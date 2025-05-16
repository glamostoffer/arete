import { create } from 'zustand/react'
import ky from 'ky'
import { AuthPaths } from '../api/urls'
import {
  getCourseCategoriesRequest,
  getCoursesRequest,
  getLessonDetailsRequest,
  getLessonsRequest,
  getUserInfoRequest,
} from './requests'
import { createJSONStorage, persist } from 'zustand/middleware'

export const useAuthStore = create(
  persist(
    (set, get) => ({
      accessToken: '',
      refreshToken: '',
      setRefreshToken: (token) => set({ refreshToken: token }),
      setAccessToken: (token) => set({ accessToken: token }),
    }),
    { name: 'authStore', storage: createJSONStorage(() => sessionStorage) }
  )
)

export const useSignUpStore = create((set, get) => ({
  email: '',
  resendCooldown: 0,
  setEmail: (email) => set({ email }),
  setResendCooldown: (resendCooldown) => set({ resendCooldown }),

  signIn: async (email, password) =>
    await ky
      .post(`${AuthPaths.SIGN_IN}`, {
        json: { login: email, password, ip: '127.0.0.1', userAgent: 'Firefox/5.1' },
      })
      .json()
      .then((res) => {
        useAuthStore.getState().setAccessToken(res.accessToken)
        useAuthStore.getState().setRefreshToken(res.refreshToken)
      }),

  signUpStart: async (login, email, password, passwordConfirmation) =>
    await ky
      .post(`${AuthPaths.SIGN_UP_START}`, {
        json: { login, email, password, passwordConfirmation },
      })
      .json()
      .then((res) => {
        get().setResendCooldown(res.resendCooldown)
      }),

  signUpFinalize: async (confirmationCode) =>
    await ky
      .post(`${AuthPaths.SIGN_UP_FINALIZE}`, {
        json: { email: get().email, confirmationCode },
      })
      .json()
      .then((res) => {
        useAuthStore.getState().setAccessToken(res.accessToken)
        useAuthStore.getState().setRefreshToken(res.refreshToken)
      }),
}))

export const useUserStore = create((set, get) => ({
  id: '',
  login: '',
  registrationDate: '',
  email: '',
  setUser: (user) => set(user),

  getUserInfo: () => getUserInfoRequest().then((res) => get().setUser(res)),
}))

export const useLearnCategoryStore = create((set, get) => ({
  categories: [],
  setCategories: (categories) => set({ categories }),
  getCategories: () =>
    getCourseCategoriesRequest().then((res) => get().setCategories(res.items)),
}))

export const useLearnCourseStore = create((set, get) => ({
  courses: [],
  setCourses: (courses) => set({ courses }),
  getCourses: (categories) =>
    getCoursesRequest(categories).then((res) => get().setCourses(res.items)),
}))

export const useLearnLessonStore = create((set, get) => ({
  lessons: [],
  setLessons: (lessons) => set({ lessons }),
  getLessons: (courseId) =>
    getLessonsRequest(courseId).then((res) => get().setLessons(res)),
}))

export const useLessonDetailsStore = create((set, get) => ({
  lesson: {},
  setLesson: (lesson) => set({ lesson }),
  getLessonDetails: (lessonId) =>
    getLessonDetailsRequest(lessonId).then((res) => get().setLesson(res)),
}))
