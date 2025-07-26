# AI Code Reviewer

It uses AI to review your code. It would fetch all code in the current branch, and evaluate your latest commit on the current branch. It would return review results to you in the terminal.

Features:

- When loading branch code, ignore files specified .gitignore.
- Avoid sending all files in every request. Let AI remember the sesson and use files updated instead.
- Test to use local LLM
- Use better prompt for better result -- maybe specify languages in question?
