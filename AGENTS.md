```markdown
# AGENTS.md File Guidelines

These guidelines are designed to ensure the consistent, efficient, and maintainable development of the AGENTS repository. Adherence to these principles is crucial for creating a robust and scalable AI agent system.

## 1. DRY (Don't Repeat Yourself)

*   **Single Responsibility Principle:** Each agent should have a single, well-defined purpose.  Avoid creating modules with overlapping functionality.
*   **Abstraction:** Use abstract classes and interfaces to decouple components and promote reusability.
*   **Standardized Components:** Define and document standard components and their interactions.
*   **Parameterization:**  Design agents to be easily parameterized with minimal changes.

## 2. KISS (Keep It Simple, Stupid)

*   **Minimalist Code:** Strive for the shortest possible code to achieve the desired functionality.  Avoid unnecessary complexity.
*   **Readability:** Prioritize code that is easy to understand and maintain. Use clear variable names, comments, and formatting.
*   **Single Responsibility:**  As mentioned above, break down complex logic into smaller, focused functions/modules.

## 3. SOLID Principles

*   **Single Responsibility:**  (Reinforced) Each class/module should have a single, well-defined responsibility.
*   **Open/Closed Principle:**  The system should be extensible through new classes/modules without modifying existing code.  (Implementation should be open for modification, but not for extension.)
*   **Liskov Substitution Principle:**  Subclasses should be substitutable for their base classes without changing the correctness of the program.
*   **Interface Segregation Principle:**  Clients shouldn't be forced to implement interfaces they don't need.
*   **Dependency Inversion Principle:**  High-level modules (classes) should not depend on low-level modules. Interfaces should define contracts.

## 4. YAGNI (You Aren't Gonna Need It)

*   **Avoid Unnecessary Code:**  Do not write code that you are not currently planning to use. Focus on the tasks needed for the current sprint/iteration.
*   **Future-Proofing:**  Design for potential future requirements without prematurely implementing them.

## 5. Development Guidelines

*   **Code Style:**  Follow a consistent code style (e.g., PEP 8 for Python) throughout the codebase.
*   **Naming Conventions:**  Adopt a consistent naming convention for variables, functions, and classes (e.g., snake_case).
*   **Comments:**  Provide clear and concise comments to explain complex logic, design decisions, and potential pitfalls.  Comments should enhance understanding, not simply restate code.
*   **Error Handling:**  Implement basic error handling to prevent crashes and provide informative error messages.  Logging is preferred.
*   **Testing:**  All development must be productive and lead to testing.  Use mocks and stubs for testing during development.
*   **Documentation:**  Document API endpoints, data structures, and any complex logic used.

## 6. File Length Constraints

*   **Maximum File Length:** 180 lines of code (excluding comments & docstrings)
*   **Code Complexity:**  Aim for a complexity score of less than 10 (measured by the number of lines of code and the depth of nesting).  A complexity score below 4 is recommended.

## 7. Test Coverage Requirements

*   **Minimum Test Coverage:** 80%
*   **Test Types:**  Include unit tests, integration tests, and end-to-end tests where applicable.
*   **Test Focus:** Tests should validate critical functionality and corner cases.

## 8.  Reporting and Review Process

*   **Regular Reviews:** Conduct weekly code reviews to identify potential issues and ensure adherence to guidelines.
*   **Documentation Updates:**  Update documentation whenever code changes.
*   **Code Quality Metrics:** Use automated code quality tools to track code metrics (e.g., cyclomatic complexity, lines of code).

## 9.  Specific File Format (Example - Python)

*   All Python files will adhere to PEP 8 style guide.
*   Docstrings will be included for all functions and classes, explaining their purpose, parameters, and return values.
*   Code will be structured using classes and modules for organization.

## 10.  Continuous Improvement

*   Regularly analyze the repository’s code quality and identify areas for improvement.
*   Implement feedback from code reviews and testing to enhance the codebase.
```