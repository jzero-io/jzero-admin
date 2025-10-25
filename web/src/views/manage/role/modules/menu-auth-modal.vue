<script setup lang="ts">
import { computed, ref, shallowRef, watch } from 'vue';
import { useLoading } from '@sa/hooks';
import { $t } from '@/locales';
import { GetAllPages, GetMenuTree, GetRoleHome, GetRoleMenus, SetRoleMenus, UpdateRoleHome } from '@/service/api';

defineOptions({
  name: 'MenuAuthModal'
});

interface Props {
  /** the roleId */
  roleUuid: string;
}

const props = defineProps<Props>();

const {
  loading: getTreeDataLoading,
  startLoading: getTreeDataStartLoading,
  endLoading: getTreeDataEndLoading
} = useLoading();

const {
  loading: setMenusConfirmLoading,
  startLoading: setMenusConfirmStartLoading,
  endLoading: setMenusConfirmEndLoading
} = useLoading();

const updateHomeLoadingRef = ref(false);

const visible = defineModel<boolean>('visible', {
  default: false
});

function closeModal() {
  visible.value = false;
}

const title = computed(() => $t('common.edit') + $t('page.manage.role.menuAuth'));

const home = shallowRef('');

async function getHome() {
  const { error, data } = await GetRoleHome(props.roleUuid);
  if (!error) {
    home.value = data;
  }
}

async function updateHome(val: string) {
  // request
  const req: Api.Manage.UpdateRoleHomeRequest = {
    roleUuid: props.roleUuid,
    home: val
  };
  updateHomeLoadingRef.value = true;
  const { error } = await UpdateRoleHome(req);
  updateHomeLoadingRef.value = false;
  if (error) return;
  home.value = val;
  window.$message?.success($t('common.editHomeSuccess'));
}

const pages = shallowRef<string[]>([]);

async function getPages() {
  const { error, data } = await GetAllPages(props.roleUuid);

  if (!error) {
    pages.value = data;
  }
}

const pageSelectOptions = computed(() => {
  const opts: CommonType.Option[] = pages.value.map(page => ({
    label: page,
    value: page
  }));

  return opts;
});

const tree = shallowRef<Api.Manage.MenuTree[]>([]);
const checks = shallowRef<string[]>([]);

function updateChecks() {
  // 定义递归检查函数
  function checkParent(node: Api.Manage.MenuTree) {
    if (node.children && node.children.length > 0) {
      const hasCheckedChildren = node.children.some(child => checks.value.includes(child.uuid));

      // 如果有任意子节点被选中，则选中父节点
      if (hasCheckedChildren && !checks.value.includes(node.uuid)) {
        checks.value.push(node.uuid);
      }
      // 如果没有子节点被选中，并且父节点当前被选中，则取消选中父节点
      else if (!hasCheckedChildren && checks.value.includes(node.uuid)) {
        checks.value.splice(checks.value.indexOf(node.uuid), 1);
      }

      // 对每个子节点递归调用 checkParent 函数
      node.children.forEach(child => {
        if (child.children && child.children.length > 0) {
          checkParent(child);
        }
      });
    }
  }

  // 遍历顶层节点并开始检查
  tree.value.forEach(topLevelNode => {
    checkParent(topLevelNode);
  });
}

async function getTree() {
  getTreeDataStartLoading();
  const { error, data } = await GetMenuTree();
  if (!error) {
    tree.value = data;
  }

  const getRoleMenusRequest: Api.Manage.GetRoleMenusRequest = {
    roleUuid: props.roleUuid
  };
  const { error: roleMenusError, data: roleMenusData } = await GetRoleMenus(getRoleMenusRequest);
  getTreeDataEndLoading();
  if (!roleMenusError) {
    checks.value = roleMenusData.menuUuids;
  }

  // 定义递归检查函数
  function checkParent(node: Api.Manage.MenuTree) {
    if (node.children && node.children.length > 0) {
      const allChildrenChecked = node.children.every(child => checks.value.includes(child.uuid));

      // 如果所有子节点都被选中，则选中父节点（如果还没有选中的话）
      if (allChildrenChecked && !checks.value.includes(node.uuid)) {
        checks.value.push(node.uuid);
      }
      // 如果不是所有子节点都被选中，并且父节点当前被选中，则取消选中父节点
      else if (!allChildrenChecked && checks.value.includes(node.uuid)) {
        checks.value.splice(checks.value.indexOf(node.uuid), 1);
      }

      // 对每个子节点递归调用 checkParent 函数
      node.children.forEach(child => checkParent(child));
    }
  }

  // 遍历顶层节点并开始检查
  tree.value.forEach(topLevelNode => {
    checkParent(topLevelNode);
  });
}

async function handleSubmit() {
  updateChecks();
  // request
  const setRoleMenusRequest: Api.Manage.SetRoleMenusRequest = {
    roleUuid: props.roleUuid,
    menuUuids: checks.value.filter(uuid => uuid !== '')
  };
  setMenusConfirmStartLoading();
  const { error } = await SetRoleMenus(setRoleMenusRequest);
  setMenusConfirmEndLoading();
  if (!error) {
    window.$message?.success?.($t('common.modifySuccess'));
    closeModal();
  }
}

function init() {
  getHome();
  getPages();
  getTree();
}

watch(visible, val => {
  if (val) {
    init();
  }
});
</script>

<template>
  <NModal v-model:show="visible" :loading="true" :title="title" preset="card" class="w-480px">
    <div class="flex-y-center gap-16px pb-12px">
      <div>{{ $t('page.manage.menu.home') }}</div>
      <NSelect :value="home" :options="pageSelectOptions" size="small" class="w-160px" @update:value="updateHome" />
      <template v-if="updateHomeLoadingRef">
        <NSpin size="small" />
      </template>
    </div>
    <template v-if="getTreeDataLoading">
      <NSpace class="h-280px" justify="center" align="center">
        <NSpin size="small" />
      </NSpace>
    </template>
    <template v-else>
      <NTree
        v-model:checked-keys="checks"
        :data="tree"
        key-field="uuid"
        checkable
        expand-on-click
        virtual-scroll
        block-line
        cascade
        class="h-280px"
      />
    </template>
    <template #footer>
      <NSpace justify="end">
        <NButton size="small" class="mt-16px" @click="closeModal">
          {{ $t('common.cancel') }}
        </NButton>
        <NButton :loading="setMenusConfirmLoading" type="primary" size="small" class="mt-16px" @click="handleSubmit">
          {{ $t('common.confirm') }}
        </NButton>
      </NSpace>
    </template>
  </NModal>
</template>

<style scoped></style>
