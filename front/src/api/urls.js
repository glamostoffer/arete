export const API_BASE_URL = 'http://localhost:8090/api/v1'

export const AuthPaths = {
  SIGN_UP_START: `${API_BASE_URL}/auth/sign-up/start`,
  SIGN_UP_FINALIZE: `${API_BASE_URL}/auth/sign-up/finalize`,
  SIGN_IN: `${API_BASE_URL}/auth/sign-in`,
  SESSION_REFRESH: `${API_BASE_URL}/auth/session/refresh`,
}

export const UserPaths = {
  GET_USER_INFO: `${API_BASE_URL}/auth/user`,
}

export const LearnPaths = {
  GET_COURSES: `${API_BASE_URL}/learn/course`,
  GET_COURSE_CATEGORIES: `${API_BASE_URL}/learn/course/categories`,
  GET_COURSE_LESSONS: `${API_BASE_URL}/learn/lesson`,
  GET_LESSON_DETAILS: `${API_BASE_URL}/learn/lesson/details`,
}
