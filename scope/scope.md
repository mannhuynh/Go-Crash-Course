# Scopes in GO

### Package:

- Package scope: all variables or functions need a key word such as `var` or `func`
- Constant variable always need to initialize the value (for both package or block scope)

### Block:

- Block level is inside a function i.e. `main()`
- Variable declaration inside a block level doesn't need the `var` keyword

### Module:

- Module is initialized with
  ```shell
  $go mod init myapp
  ```
- Now we can start create new package (folder) and create new file in that packge
- To export variable or function from new package, we need Capitalize the first letter of the varible name or function name. i.e. `MyPackageVar` or `MyPackageFunc`
- Import package name follow module name:
  ```go
  import "myapp/packagename"
  ```
- Call the variable or function using dot notation after package name. I.e. `packagename.MyPackageVar`
