# Go HTTP Server with HTML Form Handling

This Go program demonstrates a basic HTTP server that handles an HTML form submission. Upon accessing the root URL (`/`), it serves an HTML form to the client. When the form is submitted, it retrieves the form data (name, email, password) and displays it back to the user on the same page.

## Usage

1. Clone the repository or create a Go file with the provided code snippet.

2. Run the server by executing the Go file:

   ```bash
   go run main.go
   ```

3. Access the server in a web browser at `http://localhost:8080`.

## Code Explanation

### `UserInfo` struct

- `UserInfo` defines the structure for storing user information, including `Name`, `Email`, and `Password`.

### `handelIndex` function

- `handelIndex` is an HTTP handler function that manages the root URL (`/`). It does the following:
  - Parses an HTML template file named `index.html`.
  - If the request method is not POST, it executes the template to display the form.
  - If the request method is POST, it retrieves the form values (name, email, password) and displays them back to the user.

### `main` function

- The `main` function sets up the HTTP server:
  - Handles the root URL with the `handelIndex` function.
  - Serves static assets (e.g., CSS, images) from the `assets` directory.
  - Starts the HTTP server on port `8080`.

## File Structure

- `template/index.html`: Contains the HTML form template.
- `assets/`: Directory containing static files like CSS, images, etc.

## Accessing the Form

1. Access the server in a web browser at `http://localhost:8080`.
2. Fill in the form fields (name, email, password) and submit the form.
