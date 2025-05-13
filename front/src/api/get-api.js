import ky from 'ky'
import { useAuthStore } from '../model/store'
import { AuthPaths } from './urls'

export const getApi = (accessToken, refreshToken) => {
  return ky.extend({
    hooks: {
      beforeRequest: [
        (request) => {
          request.headers.set('X-Access-Token', accessToken)
        },
      ],
      afterResponse: [
        //  retry with a fresh token on a 418 error
        async (request, options, response) => {
          if (response.status === 418) {
            // Get a fresh token
            const token = await ky
              .post(`${AuthPaths.SESSION_REFRESH}`, { json: { refreshToken } })
              .json()
              .then((res) => useAuthStore.getState().setAccessToken(res.accessToken))

            // Retry with the token
            request.headers.set('X-Access-Token', token)

            return ky(request)
          }
        },
      ],
    },
  })
}
