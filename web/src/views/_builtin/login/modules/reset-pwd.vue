<script setup lang="ts">
import { computed, reactive } from 'vue';
import { useLoading } from '@sa/hooks';
import { $t } from '@/locales';
import { useRouterPush } from '@/hooks/common/router';
import { useCaptcha } from '@/hooks/business/captcha';
import { useFormRules, useNaiveForm } from '@/hooks/common/form';
import { ResetPassword } from '@/service/api';

defineOptions({
  name: 'ResetPwd'
});

const { toggleLoginModule } = useRouterPush();
const { formRef, validate } = useNaiveForm();
const { label, isCounting, loading, getCaptcha, verificationUuid } = useCaptcha();
const { loading: confirmLoading, startLoading: confirmStartLoding, endLoading: confirmEndLoading } = useLoading();

interface FormModel {
  email: string;
  code: string;
  password: string;
  confirmPassword: string;
}

const model: FormModel = reactive({
  email: '',
  code: '',
  password: '',
  confirmPassword: ''
});

type RuleRecord = Partial<Record<keyof FormModel, App.Global.FormRule[]>>;

const rules = computed<RuleRecord>(() => {
  const { formRules, createConfirmPwdRule } = useFormRules();

  return {
    phone: formRules.phone,
    password: formRules.pwd,
    confirmPassword: createConfirmPwdRule(model.password)
  };
});

async function handleSubmit() {
  await validate();
  const resetPasswordData: Api.Auth.ResetPasswordRequest = {
    verificationCode: model.code,
    email: model.email,
    verificationUuid: verificationUuid.value,
    password: model.password
  };
  confirmStartLoding();
  const { error } = await ResetPassword(resetPasswordData);
  if (!error) {
    window.$message?.success($t('common.modifySuccess'));
    await toggleLoginModule('pwd-login');
  }
  confirmEndLoading();
}
</script>

<template>
  <NForm ref="formRef" :model="model" :rules="rules" size="large" :show-label="false" @keyup.enter="handleSubmit">
    <NFormItem path="email">
      <NInput v-model:value="model.email" :placeholder="$t('page.login.common.emailPlaceholder')" />
    </NFormItem>
    <NFormItem path="code">
      <div class="w-full flex-y-center gap-16px">
        <NInput v-model:value="model.code" :placeholder="$t('page.login.common.codePlaceholder')" />
        <NButton size="large" :disabled="isCounting" :loading="loading" @click="getCaptcha('email', model.email)">
          {{ label }}
        </NButton>
      </div>
    </NFormItem>
    <NFormItem path="password">
      <NInput
        v-model:value="model.password"
        type="password"
        show-password-on="click"
        :placeholder="$t('page.login.common.passwordPlaceholder')"
      />
    </NFormItem>
    <NFormItem path="confirmPassword">
      <NInput
        v-model:value="model.confirmPassword"
        type="password"
        show-password-on="click"
        :placeholder="$t('page.login.common.confirmPasswordPlaceholder')"
      />
    </NFormItem>
    <NSpace vertical :size="18" class="w-full">
      <NButton type="primary" size="large" round block :loading="confirmLoading" @click="handleSubmit">
        {{ $t('common.confirm') }}
      </NButton>
      <NButton size="large" round block @click="toggleLoginModule('pwd-login')">
        {{ $t('page.login.common.back') }}
      </NButton>
    </NSpace>
  </NForm>
</template>

<style scoped></style>
