# Task Manager CLI

## 📖 Overview

This is a simple command-line interface (CLI) task manager written in Go. It allows users to manage their daily tasks quickly and efficiently directly from the terminal. Tasks are stored in a JSON file (`data.json`), making the data portable and easy to read.

## 🚀 How to Use

### 🛠️ Available Commands

* **`tkm add <description>`** or **`tkm a <description>`**: Adds a new task to your list. The description must be provided in double quotes if it contains spaces. The task is created with a "todo" status.
* **`tkm list`** or **`tkm l`**: Lists all tasks.
    * Flags to filter the list:
        * `--todo`: Lists only tasks with "todo" status.
        * `--done`: Lists only tasks with "done" status.
        * `--in-progress`: Lists only tasks with "in-progress" status.
* **`tkm update <id> <new_description>`** or **`tkm u <id> <new_description>`**: Updates the description of a task based on its ID.
* **`tkm mark <id> --done`** or **`tkm m <id> --done`**: Marks a task as "done".
* **`tkm mark <id> --in-progress`** or **`tkm m <id> --in-progress`**: Marks a task as "in-progress".
* **`tkm delete <id>`** or **`tkm d <id>`**: Deletes a task based on its ID.

---

### 🏃 Examples

Here are some examples of how to run the commands:

1.  **Adding a new task:**
    ```sh
    tkm add "Study for the Go exam"
    ```

2.  **Listing all tasks:**
    ```sh
    tkm list
    ```

3.  **Listing only 'todo' tasks:**
    ```sh
    tkm list --todo
    ```

4.  **Marking a task as 'done':**
    ```sh
    tkm mark 1 --done
    ```

5.  **Updating a task's description:**
    ```sh
    tkm update 1 "Study Go for the final project"
    ```

6.  **Deleting a task:**
    ```sh
    tkm delete 2
    ```

---

## 🔧 Installation

To install this CLI, you must have **Go** installed on your system.

1.  Clone the repository to your local machine:
    ```sh
    git clone <URL_do_seu_repositório>
    cd <nome_do_seu_repositório>
    ```

2.  Build the executable file:
    ```sh
    go build -o tkm
    ```

3.  (Optional) Move the executable to a directory in your system's `PATH` to run the command from any location:
    ```sh
    # For Linux/macOS
    sudo mv tkm /usr/local/bin/

    # For Windows, move the .exe file to a directory in your PATH manually.
    ```

---

## 📁 Data Structure

Tasks are stored in a JSON file in the following format:

```json
[
  {
    "id": 1,
    "desc": "Study Go for the project",
    "status": "in-progress",
    "createdAt": "2025-08-06T21:57:05.123456Z",
    "updateAt": "2025-08-07T09:28:05.123456Z"
  },
  {
    "id": 2,
    "desc": "Go grocery shopping",
    "status": "todo",
    "createdAt": "2025-08-07T09:28:05.123456Z"
  }
]
