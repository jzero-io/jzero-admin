import { computed, ref } from 'vue';
import { useCountDown, useLoading } from '@sa/hooks';
import { $t } from '@/locales';
import { REG_EMAIL, REG_PHONE } from '@/constants/reg';
import { SendVerificationCode } from '@/service/api';

export function useCaptcha() {
  const { loading, startLoading, endLoading } = useLoading();
  const { count, start, stop, isCounting } = useCountDown(30);
  const verificationUuid = ref(''); // 使用 ref 来存储 UUID

  const label = computed(() => {
    let text = $t('page.login.codeLogin.getCode');

    const countingLabel = $t('page.login.codeLogin.reGetCode', { time: count.value });

    if (loading.value) {
      text = '';
    }

    if (isCounting.value) {
      text = countingLabel;
    }

    return text;
  });

  function isEmailValid(email: string) {
    if (email.trim() === '') {
      window.$message?.error?.($t('form.email.required'));

      return false;
    }

    if (!REG_EMAIL.test(email)) {
      window.$message?.error?.($t('form.email.invalid'));

      return false;
    }

    return true;
  }

  function isPhoneValid(phone: string) {
    if (phone.trim() === '') {
      window.$message?.error?.($t('form.phone.required'));

      return false;
    }

    if (!REG_PHONE.test(phone)) {
      window.$message?.error?.($t('form.phone.invalid'));

      return false;
    }

    return true;
  }

  async function getCaptcha(type: string, value: string) {
    let valid;
    if (type === 'email') {
      valid = isEmailValid(value);
    } else {
      valid = isPhoneValid(value);
    }

    if (!valid || loading.value) {
      return;
    }

    startLoading();

    try {
      const params: Api.Auth.SendVerificationCodeRequest = {
        verificationType: type,
        email: value
      };

      // 发送请求
      const { data, error } = await SendVerificationCode(params);

      if (!error) {
        window.$message?.success?.($t('page.login.codeLogin.sendCodeSuccess'));
        verificationUuid.value = data?.verificationUuid;
      }
    } finally {
      start();

      endLoading();
    }
  }

  return {
    label,
    start,
    stop,
    isCounting,
    loading,
    getCaptcha,
    verificationUuid
  };
}
