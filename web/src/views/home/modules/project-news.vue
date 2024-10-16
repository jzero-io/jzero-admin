<script setup lang="ts">
import { computed } from 'vue';
import { $t } from '@/locales';

defineOptions({
  name: 'ProjectNews'
});

interface NewsItem {
  id: number;
  content: string;
  time: string;
}

const newses = computed<NewsItem[]>(() => [
  { id: 1, content: $t('page.home.projectNews.desc1'), time: '2024-04-07' },
  { id: 2, content: $t('page.home.projectNews.desc2'), time: '2024-05-08' },
  { id: 3, content: $t('page.home.projectNews.desc3'), time: '2024-05-23' },
  { id: 4, content: $t('page.home.projectNews.desc4'), time: '2024-06-03' },
  { id: 5, content: $t('page.home.projectNews.desc5'), time: '2024-10-14' }
]);
</script>

<template>
  <NCard :title="$t('page.home.projectNews.title')" :bordered="false" size="small" segmented class="card-wrapper">
    <template #header-extra>
      <a class="text-primary" href="https://github.com/jzero-io/jzero" target="_blank">{{ $t('page.home.projectNews.moreNews') }}</a>
    </template>
    <NList>
      <NListItem v-for="item in newses" :key="item.id">
        <template #prefix>
          <SoybeanAvatar class="size-48px!" />
        </template>
        <NThing :title="item.content" :description="item.time" />
      </NListItem>
    </NList>
  </NCard>
</template>

<style scoped></style>
