# Lesson 01 — App Router

## What it is

The **App Router** is the current Next.js routing system. Routes come from the `app/` directory (often under `src/app/`).

Each folder is a URL segment. Special filenames define behavior:

| File | Role |
| :--- | :--- |
| `page.tsx` | UI for that route |
| `layout.tsx` | Shared layout (wraps pages) |
| `loading.tsx` | Loading UI |
| `error.tsx` | Error UI |
| `not-found.tsx` | 404 UI |
| `favicon.ico` | Favicon (file convention at `app/` root) |

Example:

```
src/app/
  page.tsx              → /
  dashboard/
    page.tsx            → /dashboard
  layout.tsx            → root layout for all routes
```

## App Router vs Pages Router

| | App Router | Pages Router |
| :--- | :--- | :--- |
| Folder | `app/` | `pages/` |
| Status | Current default | Legacy |
| Layouts | Nested `layout.tsx` | `_app` / custom layouts |
| Server Components | First-class | Not the model |

This project chose App Router — see [ADR-010](../ADRs/010-next-app-router.md).

## Rule of thumb

- Folders under `app/` = **routes**
- Do not put random “organization” folders inside `app/` for non-route assets unless you intend a URL segment
- Shared UI and logic belong outside `app/` (`components/`, `hooks/`, `lib/`)
