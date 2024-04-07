#### 1. when using react reducer
```
Unexpected Application Error!
state.user is undefined
```
remember register store
```
export const store = configureStore({
  reducer: {
    user: userReducer
  },
});
```
#### 2. react route
```
Absolute route path "/" nested under path "/user" is not valid. An absolute child route path must start with the combined path of all its parent routes.
```
No slash in sub path
```
      <Route path="/user" element={<ProtectedLayout privilege={PRIVILEGE_USER} />}>
        <Route path="" element={<HomeUserPage />} />
      </Route>

      <Route path="/admin" element={<ProtectedLayout privilege={PRIVILEGE_ADMIN} />}>
        <Route path="" element={<HomeAdminPage />} />
        <Route path="user" element={<AdminUserPage />} />
        /*error here*/<Route path="/user" element={<AdminUserPage />} />
      </Route>
```