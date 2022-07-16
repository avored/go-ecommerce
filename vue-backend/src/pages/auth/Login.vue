<script setup lang="ts">
import axios from 'axios'
import { useRouter } from 'vue-router'
import { ref } from "vue"
import { useI18n } from 'vue-i18n'

const { t } = useI18n({
    inheritLocale: true,
    useScope: 'global'
})

interface User {
    email?: string;
    password?: string;
}

const router = useRouter()
const user = ref<User>({});

const loginFormSubmit = () => {
    const endpointBaseUrl = import.meta.env.VITE_APP_BACKEND_BASE_URL

    axios({
        method: 'POST',
        url: endpointBaseUrl + "/admin/login",
        data: user.value,
    }).then(response => {
            router.push({ name: 'dashboard' })
        })
        .catch(error => {
            console.error("There was an error!", error);
        });
};
</script>

<template>
    <div class="h-screen min-h-full flex items-center justify-center py-12  px-4 sm:px-6 lg:px-8">
        <div class="max-w-md w-full shadow p-8">
            <div>
                <img  class="mx-auto h-12 w-auto" src="@/assets/logo.svg" alt="AvoRed E commerce" />
                <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
                    {{ t('sign_into_avored_account') }}
                </h2>
            </div>
            <form class="mt-8 space-y-6" @submit.prevent="loginFormSubmit">
                <input type="hidden" name="remember" value="true" />
                <div class="rounded-md shadow-sm">
                    <div class="mt-5">
                        <label for="email-address" class="block text-gray-500">
                            {{ t('email_address') }}
                        </label>
                        <input id="email-address" type="email" 
                            autocomplete="email" 
                            v-model="user.email"
                            required class="block w-full rounded text-sm text-gray-800" :placeholder="t('email_address')" />
                    </div>
                    <div class="mt-5">
                        <label for="password" class="block text-gray-500">
                            {{ t('password') }}
                        </label>
                        <input id="password" type="password"  required
                            v-model="user.password" 
                            class="block w-full rounded text-sm text-gray-800" 
                            :placeholder="t('password')" />
                    </div>
                </div>

                <div class="flex items-center justify-between">
                    <div class="flex items-center">
                        <input
                        id="remember-me"
                        name="remember-me"
                        type="checkbox"
                        class="
                            h-4
                            w-4
                            text-red-600
                            focus:ring-red-500
                            border-gray-300
                            rounded
                        "
                        />
                        <label for="remember-me" class="ml-2 block text-sm text-gray-900">
                            {{ t('remember_me') }}
                        </label>
                    </div>
                    <div class="text-sm">
                        <a href="#" class="font-medium text-red-600 hover:text-red-500">
                            {{ t('forgot_password') }}
                        </a>
                    </div>
                </div>

                <div>
                    <button
                        type="submit"
                        class="
                        group
                        relative
                        w-full
                        flex
                        justify-center
                        py-2
                        px-4
                        border border-transparent
                        text-sm
                        font-medium
                        rounded-md
                        text-white
                        bg-red-600
                        hover:bg-red-700
                        focus:outline-none
                        focus:ring-2
                        focus:ring-offset-2
                        focus:ring-red-500
                        "
                    >
                        <span class="absolute left-0 inset-y-0 flex items-center pl-3">
                        <!-- Heroicon name: solid/lock-closed -->
                        <svg
                            class="h-5 w-5 text-red-500 group-hover:text-red-400"
                            xmlns="http://www.w3.org/2000/svg"
                            viewBox="0 0 20 20"
                            fill="currentColor"
                            aria-hidden="true"
                        >
                            <path
                            fill-rule="evenodd"
                            d="M5 9V7a5 5 0 0110 0v2a2 2 0 012 2v5a2 2 0 01-2 2H5a2 2 0 01-2-2v-5a2 2 0 012-2zm8-2v2H7V7a3 3 0 016 0z"
                            clip-rule="evenodd"
                            />
                        </svg>
                        </span>
                            {{  t('sign_in') }}
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>
