# Task Manager CLI

## 📖 Overview

This is a simple command-line interface (CLI) task manager written in Go. It allows users to manage their daily tasks quickly and efficiently directly from the terminal. Tasks are stored in a JSON file (`data.json`), making the data portable and easy to read.

## 🚀 How to Use

### 🛠️ Available Commands

  * `tkm add <description>` or `tkm a <description>`: Adds a new task to your list. The description must be provided in double quotes if it contains spaces. The task is created with a "todo" status.
  * `tkm list` or `tkm l`: Lists all tasks.
      * Flags to filter the list:
          * `--todo`: Lists only tasks with "todo" status.
          * `--done`: Lists only tasks with "done" status.
          * `--in-progress`: Lists only tasks with "in-progress" status.
  * `tkm update <id> <new_description>` or `tkm u <id> <new_description>`: Updates the description of a task based on its ID.
  * `tkm mark <id> --done` or `tkm m <id> --done`: Marks a task as "done".
  * `tkm mark <id> --in-progress` or `tkm m <id> --in-progress`: Marks a task as "in-progress".
  * `tkm delete <id>` or `tkm d <id>`: Deletes a task based on its ID.

## 📁 Data Structure

Tasks are stored in a JSON file in the following format:

```json
[
  {
    "id": 1,
    "desc": "Estudar Go para o projeto",
    "status": "in-progress",
    "createdAt": "2025-08-06T21:57:05.123456Z",
    "updateAt": "2025-08-07T09:28:05.123456Z"
  },
  {
    "id": 2,
    "desc": "Fazer compras de supermercado",
    "status": "todo",
    "createdAt": "2025-08-07T09:28:05.123456Z"
  }
]
```
