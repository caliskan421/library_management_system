# Hook: Post-Edit

## Description
Verification steps after editing Go source files.

## After Editing Model Files (`internal/model/*.go`)
- Check that Bun struct tags are correct
- Verify JSON tags match API spec (camelCase)
- Ensure `bun:"table:name,alias:x"` is set
- Check if database migration needs update

## After Editing Handler Files (`internal/handler/*.go`)
- Verify error response format matches `dto.ErrorResponse`
- Check HTTP status codes match API spec
- Ensure `c.Locals()` type assertions are safe

## After Editing Service Files (`internal/service/*.go`)
- Verify domain errors are properly defined
- Check business rule implementation matches requirements
- Ensure repository calls use correct context

## After Editing Repository Files (`internal/repository/*.go`)
- Verify SQL queries use parameterized values (no string concat)
- Check pagination offset calculation: `(page - 1) * limit`
- Ensure proper Bun query builder usage

## Token Optimization
- Read only the changed file, not the entire package
- Cross-reference with dto/ only if request/response shape changed
- Skip reading CLAUDE.md if working within a single layer
