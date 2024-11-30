<script setup lang="ts">
import { computed, shallowRef, watch } from 'vue';
import { useLoading } from '@sa/hooks';
import { $t } from '@/locales';
import { GetAllPages, GetMenuTree, GetRoleMenus, SetRoleMenus } from '@/service/api';

defineOptions({
  name: 'MenuAuthModal'
});

interface Props {
  /** the roleId */
  roleId: number;
}

const props = defineProps<Props>();
const {
  loading: getTreeDataLoading,
  startLoading: getTreeDataStartLoading,
  endLoading: getTreeDataEndLoading
} = useLoading();
const {
  loading: setMenusConfirmLoding,
  startLoading: setMenusConfirmStartLoading,
  endLoading: setMenusConfirmEndLoading
} = useLoading();

const visible = defineModel<boolean>('visible', {
  default: false
});

function closeModal() {
  visible.value = false;
}

const title = computed(() => $t('common.edit') + $t('page.manage.role.menuAuth'));

const home = shallowRef('');

async function getHome() {
  home.value = 'home';
}

async function updateHome(val: string) {
  // request

  home.value = val;
}

const pages = shallowRef<string[]>([]);

async function getPages() {
  const { error, data } = await GetAllPages();

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
const checks = shallowRef<number[]>([]);

async function getTree() {
  getTreeDataStartLoading();
  const { error, data } = await GetMenuTree();
  if (!error) {
    tree.value = data;
  }

  const getRoleMenusRequest: Api.Manage.GetRoleMenusRequest = {
    roleId: props.roleId
  };
  const { error: roleMenusError, data: roleMenusData } = await GetRoleMenus(getRoleMenusRequest);
  getTreeDataEndLoading();
  if (!roleMenusError) {
    checks.value = roleMenusData;
  }
}

async function handleSubmit() {
  // request
  const setRoleMenusRequest: Api.Manage.SetRoleMenusRequest = {
    roleId: props.roleId,
    menuIds: checks.value
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
        key-field="id"
        checkable
        expand-on-click
        virtual-scroll
        block-line
        class="h-280px"
      />
    </template>
    <template #footer>
      <NSpace justify="end">
        <NButton size="small" class="mt-16px" @click="closeModal">
          {{ $t('common.cancel') }}
        </NButton>
        <NButton :loading="setMenusConfirmLoding" type="primary" size="small" class="mt-16px" @click="handleSubmit">
          {{ $t('common.confirm') }}
        </NButton>
      </NSpace>
    </template>
  </NModal>
</template>

<style scoped></style>
