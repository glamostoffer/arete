package http

import (
	"context"

	"github.com/glamostoffer/arete/gateway/internal/service/dto"
)

type auth interface {
	StartSignUp(ctx context.Context, req dto.StartSignUpRequest) (res dto.StartSignUpResponse, err error)
	ConfirmEmail(ctx context.Context, req dto.ConfirmEmailRequest) (res dto.ConfirmEmailResponse, err error)
	SignIn(ctx context.Context, req dto.SignInRequest) (res dto.SignInResponse, err error)

	VerifyCredentials(ctx context.Context, req dto.VerifyCredentialsRequest) (res dto.VerifyCredentialsResponse, err error)
	RefreshSession(ctx context.Context, req dto.RefreshSessionRequest) (res dto.RefreshSessionResponse, err error)

	GetUserInfo(ctx context.Context, req dto.GetUserInfoRequest) (res dto.GetUserInfoResponse, err error)
}

type learning interface {
	GetCourseCategories(ctx context.Context, req dto.GetCourseCategoriesRequest) (res dto.GetCourseCategoriesResponse, err error)
	GetCourses(ctx context.Context, req dto.GetCoursesRequest) (res dto.GetCoursesResponse, err error)
	GetCourseLessons(ctx context.Context, req dto.GetCourseLessonsRequest) (res dto.GetCourseLessonsResponse, err error)
	GetLessonDetails(ctx context.Context, req dto.GetLessonDetailsRequest) (res dto.GetLessonDetailsResponse, err error)
}
