<script setup lang="ts">
import { computed } from 'vue';
import { $t } from '@/locales';
import { useAppStore } from '@/store/modules/app';
import { useAuthStore } from '@/store/modules/auth';

defineOptions({
  name: 'HeaderBanner'
});

const appStore = useAppStore();
const authStore = useAuthStore();

const gap = computed(() => (appStore.isMobile ? 0 : 16));

</script>

<template>
  <NCard :bordered="false" class="card-wrapper">
    <NGrid :x-gap="gap" :y-gap="16" responsive="screen" item-responsive>
      <NGi span="24 s:24 m:18">
        <div class="flex-y-center">
          <div class="w-72px shrink-0 overflow-hidden rd-1/2">
            <img src="@/assets/imgs/jzero.svg" class="w-full h-auto" />
          </div>
          <div class="pl-12px">
            <h3 class="text-18px font-semibold">
              {{ $t('page.home.greeting', { username: authStore.userInfo.username }) }}
            </h3>
            <p class="text-#999 leading-30px">{{ $t('page.home.weatherDesc') }}</p>
          </div>
        </div>
      </NGi>
    </NGrid>
  </NCard>
</template>

<style scoped></style>
