/*
Middleware RBAC *

Вам необходимо создать Middleware
RoleBasedAuthMiddleware(allowedRoles []string, next http.Handler) http.Handler
для HTTP-сервера, который будет ограничивать доступ к определенным ресурсам
в зависимости от роли пользователя.

Middleware должен проверять роль пользователя и разрешать или запрещать
доступ к ресурсам.

Роль передается в заголовке X-User-Role

Нужно написать два обработчика
для пути "/admin"
  AdminHandler(w http.ResponseWriter, r *http.Request)
для пути "/user"
  UserHandler(w http.ResponseWriter, r *http.Request)

Не забудьте про функцию main:

func main() {
  allowedAdminRoles := []string{"admin", "superadmin"}
  allowedUserRoles := []string{"user"}

  // Создание маршрута и применение Middleware для пути "/admin".
  adminHandler := RoleBasedAuthMiddleware(allowedAdminRoles,
                                          http.HandlerFunc(AdminHandler))
  http.Handle("/admin", adminHandler)

  // Создание маршрута и применение Middleware для пути "/user".
  userHandler := RoleBasedAuthMiddleware(allowedUserRoles,
                                         http.HandlerFunc(UserHandler))
  http.Handle("/user", userHandler)

  // Запуск веб-сервера на порту 8080.
  http.ListenAndServe(":8080", nil)
}

Примечания

Что нужно добавить:
Header: X-User-Role

Ожидаемый результат: Если в запрос будет прокинут X-User-Role - нужно
выводить его в таком формате.
Если пользователь - admin:
Admin Resource
Если пользователь - user:
User Resource
Если заголовка нет или роль пользователя не подходит - нужно возвращать
http.StatusForbidden
*/

package main

import (
  "fmt"
  "slices"
  "context"
  "net/http"
)


func UserHandler(w http.ResponseWriter, r *http.Request) {
  role := r.Context().Value("Role")
  if role == nil {
    w.WriteHeader(http.StatusForbidden)
  } else {
    fmt.Fprintf(w, "User Resource\n")
  }
}

func AdminHandler(w http.ResponseWriter, r *http.Request) {
  role := r.Context().Value("Role")
  if role == nil {
    w.WriteHeader(http.StatusForbidden)
  } else {
    fmt.Fprintf(w, "Admin Resource\n")
  }
}

func RoleBasedAuthMiddleware(allowedRoles []string,
                             next http.Handler) http.Handler {
  return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request) {
    role := r.Header.Get("X-User-Role")
    var ctx context.Context
    if len(role) == 0 || !slices.Contains(allowedRoles, role) {
      ctx = r.Context()
    } else {
      ctx = context.WithValue(r.Context(), "Role", role)
    }
    next.ServeHTTP(w, r.WithContext(ctx))
  })
}

func main() {
  allowedAdminRoles := []string{"admin", "superadmin"}
  allowedUserRoles := []string{"user"}

  // Создание маршрута и применение Middleware для пути "/admin".
  adminHandler := RoleBasedAuthMiddleware(allowedAdminRoles,
                                          http.HandlerFunc(AdminHandler))
  http.Handle("/admin", adminHandler)

  // Создание маршрута и применение Middleware для пути "/user".
  userHandler := RoleBasedAuthMiddleware(allowedUserRoles,
                                         http.HandlerFunc(UserHandler))
  http.Handle("/user", userHandler)

  // Запуск веб-сервера на порту 8080.
  http.ListenAndServe(":8080", nil)
}
