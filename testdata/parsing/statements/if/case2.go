package main
func main() {
    if ; e {
        return 1
    }
}
-----
Go file
  PackageDeclaration(main)
    PsiElement(KEYWORD_PACKAGE)('package')
    PsiWhiteSpace(' ')
    PsiElement(IDENTIFIER)('main')
  PsiWhiteSpace('\n')
  FunctionDeclaration(main)
    PsiElement(KEYWORD_FUNC)('func')
    PsiWhiteSpace(' ')
    LiteralIdentifierImpl
      PsiElement(IDENTIFIER)('main')
    PsiElement(()('(')
    PsiElement())(')')
    PsiWhiteSpace(' ')
    BlockStmtImpl
      PsiElement({)('{')
      PsiWhiteSpace('\n')
      PsiWhiteSpace('    ')
      IfStmtImpl
        PsiElement(KEYWORD_IF)('if')
        PsiWhiteSpace(' ')
        EmptyStmt
          <empty list>
        PsiElement(;)(';')
        PsiWhiteSpace(' ')
        LiteralExpressionImpl
          LiteralIdentifierImpl
            PsiElement(IDENTIFIER)('e')
        PsiWhiteSpace(' ')
        BlockStmtImpl
          PsiElement({)('{')
          PsiWhiteSpace('\n')
          PsiWhiteSpace('        ')
          ReturnStmtImpl
            PsiElement(KEYWORD_RETURN)('return')
            PsiWhiteSpace(' ')
            LiteralExpressionImpl
              LiteralIntegerImpl
                PsiElement(LITERAL_INT)('1')
          PsiWhiteSpace('\n')
          PsiWhiteSpace('    ')
          PsiElement(})('}')
      PsiWhiteSpace('\n')
      PsiElement(})('}')
