import { getApi } from '../api/get-api'
import { LearnPaths, UserPaths } from '../api/urls'
import { useAuthStore } from './store'

export const getUserInfoRequest = async () => {
  const accessToken = useAuthStore.getState().accessToken
  const refreshToken = useAuthStore.getState().refreshToken
  return getApi(accessToken, refreshToken).get(`${UserPaths.GET_USER_INFO}`).json()
}

export const getCourseCategoriesRequest = async () => {
  const accessToken = useAuthStore.getState().accessToken
  const refreshToken = useAuthStore.getState().refreshToken
  return getApi(accessToken, refreshToken)
    .get(`${LearnPaths.GET_COURSE_CATEGORIES}`)
    .json()
}

export const getCoursesRequest = async (categories) => {
  const accessToken = useAuthStore.getState().accessToken
  const refreshToken = useAuthStore.getState().refreshToken
  return getApi(accessToken, refreshToken)
    .post(`${LearnPaths.GET_COURSES}`, { json: { offset: 0, limit: 100, categories } })
    .json()
}

export const getLessonsRequest = async (courseID) => {
  const accessToken = useAuthStore.getState().accessToken
  const refreshToken = useAuthStore.getState().refreshToken
  return getApi(accessToken, refreshToken)
    .post(`${LearnPaths.GET_COURSE_LESSONS}`, {
      json: { offset: 0, limit: 100, courseID },
    })
    .json()
}

export const getLessonDetailsRequest = async (lessonID) => {
  const accessToken = useAuthStore.getState().accessToken
  const refreshToken = useAuthStore.getState().refreshToken
  return getApi(accessToken, refreshToken)
    .post(`${LearnPaths.GET_LESSON_DETAILS}`, { json: { lessonID } })
    .json()
}
