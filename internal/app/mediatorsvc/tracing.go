package mediatorsvc

import (
	"context"

	"go.opencensus.io/trace"
)

// Tracing represents the decorator pattern to add tracing to services
func Tracing() Middleware {
	return func(s Service) Service {
		return &tracingMiddleware{
			service: s,
		}
	}
}

type tracingMiddleware struct {
	service Service
}

func (m *tracingMiddleware) FetchUsers(ctx context.Context, location string, months int, perPage int) error {
	ctx, span := trace.StartSpan(ctx, "FetchUsers")
	defer span.End()

	return m.service.FetchUsers(ctx, location, months, perPage)
}

func (m *tracingMiddleware) FetchRepos(ctx context.Context, userPerPage, repoPerPage int, reset bool) error {
	ctx, span := trace.StartSpan(ctx, "FetchRepos")
	defer span.End()

	return m.service.FetchRepos(ctx, userPerPage, repoPerPage, reset)
}

func (m *tracingMiddleware) UpdateUserCount(ctx context.Context) error {
	ctx, span := trace.StartSpan(ctx, "UpdateUserCount")
	defer span.End()

	return m.service.UpdateUserCount(ctx)
}

func (m *tracingMiddleware) UpdateRepoCount(ctx context.Context) error {
	ctx, span := trace.StartSpan(ctx, "UpdateRepoCount")
	defer span.End()

	return m.service.UpdateRepoCount(ctx)
}

func (m *tracingMiddleware) UpdateReposMostRecent(ctx context.Context, perPage int) error {
	ctx, span := trace.StartSpan(ctx, "UpdateReposMostRecent")
	defer span.End()

	return m.service.UpdateReposMostRecent(ctx, perPage)
}

func (m *tracingMiddleware) UpdateRepoCountByUser(ctx context.Context, perPage int) error {
	ctx, span := trace.StartSpan(ctx, "UpdateRepoCountByUser")
	defer span.End()

	return m.service.UpdateRepoCountByUser(ctx, perPage)
}

func (m *tracingMiddleware) UpdateReposMostStars(ctx context.Context, perPage int) error {
	ctx, span := trace.StartSpan(ctx, "UpdateReposMostStars")
	defer span.End()

	return m.service.UpdateReposMostStars(ctx, perPage)
}

func (m *tracingMiddleware) UpdateReposMostForks(ctx context.Context, perPage int) error {
	ctx, span := trace.StartSpan(ctx, "UpdateReposMostForks")
	defer span.End()

	return m.service.UpdateReposMostForks(ctx, perPage)
}

func (m *tracingMiddleware) UpdateLanguagesMostPopular(ctx context.Context, perPage int) error {
	ctx, span := trace.StartSpan(ctx, "UpdateLanguagesMostPopular")
	defer span.End()

	return m.service.UpdateLanguagesMostPopular(ctx, perPage)
}

func (m *tracingMiddleware) UpdateMostRecentReposByLanguage(ctx context.Context, perPage int) error {
	ctx, span := trace.StartSpan(ctx, "UpdateMostRecentReposByLanguage")
	defer span.End()

	return m.service.UpdateMostRecentReposByLanguage(ctx, perPage)
}

func (m *tracingMiddleware) UpdateReposByLanguage(ctx context.Context, perPage int) error {
	ctx, span := trace.StartSpan(ctx, "UpdateReposByLanguage")
	defer span.End()

	return m.service.UpdateReposByLanguage(ctx, perPage)
}

func (m *tracingMiddleware) UpdateProfile(ctx context.Context, numWorkers int) error {
	ctx, span := trace.StartSpan(ctx, "UpdateProfile")
	defer span.End()

	return m.service.UpdateProfile(ctx, numWorkers)
}

func (m *tracingMiddleware) UpdateMatches(ctx context.Context) error {
	ctx, span := trace.StartSpan(ctx, "UpdateMatches")
	defer span.End()

	return m.service.UpdateMatches(ctx)
}

func (m *tracingMiddleware) UpdateUsersByCompany(ctx context.Context, min, max int) error {
	ctx, span := trace.StartSpan(ctx, "UpdateUsersByCompany")
	defer span.End()

	return m.service.UpdateUsersByCompany(ctx, min, max)
}

func (m *tracingMiddleware) UpdateCompanyCount(ctx context.Context) error {
	ctx, span := trace.StartSpan(ctx, "UpdateCompanyCount")
	defer span.End()

	return m.service.UpdateCompanyCount(ctx)
}
