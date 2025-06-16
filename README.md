# soco1508/inventory

soco1508/inventory is a full-stack inventory management application designed to track products, sales, purchases, and expenses. It features a Go-based backend and a Next.js frontend with a responsive user interface, Redux for state management, and Recharts for data visualization.

## Key Features

- **Comprehensive Dashboard:** Visualizes key metrics including sales summaries, purchase summaries, expense summaries, and popular products.
- **Product Management:** Allows users to view, create, and search for products within the inventory.
- **Inventory Tracking:** Provides an overview of current stock levels and detailed product information.
- **User Management:** Lists registered users of the system.
- **Expense Management:** Tracks expenses, categorized for better financial insights, with filtering options.
- **Data Visualization:** Utilizes charts (Pie, Area, Bar) for intuitive data representation on the dashboard and expenses page.
- **Seed Data Management:** Includes a backend CLI tool for batch inserting initial data into various database tables.
- **Responsive UI:** Sidebar navigation and layout adapt to different screen sizes.
- **Theme Customization:** Supports dark and light mode themes, configurable by the user.

## Tech Stack

**Backend:**

- Go (version 1.23.2)
- Gin (HTTP web framework)
- SQLx (PostgreSQL driver extension)
- PostgreSQL (Database)
- `godotenv` (Environment variable loading)
- `envconfig` (Configuration parsing from environment variables)

**Frontend:**

- Next.js (React framework, version 14.2.28)
- React (version 18)
- TypeScript
- Redux Toolkit (State management, including RTK Query for API interaction and `redux-persist` for state persistence)
- Tailwind CSS (Utility-first CSS framework)
- Recharts (Composable charting library)
- Axios (Promise-based HTTP client for API calls, managed via RTK Query)
- Lucide Icons (Icon library)

## Getting Started

### Prerequisites

- Go (version 1.23.2 or later)
- Node.js (version 20 or later)
- npm, yarn, or pnpm
- PostgreSQL server

### Backend Setup

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/soco1508/inventory.git
    cd inventory/backend
    ```

2.  **Set up environment variables:**
    Create a `.env` file in the `backend` directory. Refer to `config/config.go` for the required environment variables.
    Example `.env` file:

    ```env
    DATABASE_HOST=localhost
    DATABASE_PORT=5432
    DATABASE_USERNAME=your_postgres_user
    DATABASE_PASSWORD=your_postgres_password
    DATABASE_NAME=inventory_db
    SERVER_HOST=0.0.0.0
    SERVER_PORT=8080
    ```

3.  **Initialize Database:**
    Ensure your PostgreSQL server is running and the database specified in the `.env` file (e.g., `inventory_db`) exists. You will need to manually create the database and then the necessary tables. The table schemas can be inferred from the model definitions in `internal/db/models/` and the seed data structure in `seedData/`.

4.  **Install Go dependencies:**

    ```bash
    go mod tidy
    ```

5.  **Run the backend server:**
    ```bash
    go run cmd/server/main.go
    ```
    The server will start on the address specified by `SERVER_HOST` and `SERVER_PORT` (default: `0.0.0.0:8080`).

### Frontend Setup

1.  **Navigate to the frontend directory:**

    ```bash
    # From the repository root
    cd frontend
    # Or, if you are in the backend directory
    # cd ../frontend
    ```

2.  **Install Node.js dependencies:**

    ```bash
    npm install
    # or
    # yarn install
    # or
    # pnpm install
    ```

3.  **Set up environment variables:**
    Create a `.env.local` file in the `frontend` directory. This file should specify the base URL for the backend API.
    Example `.env.local` file:

    ```env
    NEXT_PUBLIC_API_BASE_URL=http://localhost:8080
    ```

    Adjust the URL if your backend is running on a different host or port.

4.  **Run the frontend development server:**
    ```bash
    npm run dev
    # or
    # yarn dev
    ```
    The application will be accessible at `http://localhost:3000`.

## Backend CLI for Seed Data

The backend includes a command-line interface (CLI) tool to batch insert seed data into the database tables. This is useful for initializing your development or testing environment. The seed data is located in the `backend/seedData/` directory in JSON format.

**Usage:**
Navigate to the `backend` directory and run the CLI using `go run`:

```bash
go run cmd/cli/main.go -input <path_to_json_file> -table <TableName>
```

**Example:**
To insert user data:

```bash
go run cmd/cli/main.go -input seedData/users.json -table Users
```

Available table names for the `-table` flag (case-sensitive, defined in `pkg/db/tableName.go`):

- `Users`
- `Products`
- `Sales`
- `SalesSummary`
- `Purchases`
- `PurchaseSummary`
- `Expenses`
- `ExpenseSummary`
- `ExpenseByCategory`

Ensure your database schema is created before running the seed command for the respective tables.

## Project Structure

The repository is organized into two main directories:

- `backend/`: Contains the Go application responsible for the API and business logic.
  - `cmd/`: Entry points for the server (`cmd/server/main.go`) and the data seeding CLI (`cmd/cli/main.go`).
  - `config/`: Configuration loading logic (`config/config.go`).
  - `internal/`: Core application logic.
    - `api/`: HTTP handlers (`handler/`) and API routes (`routes/`).
    - `db/`: Database models (`models/`) and repository layer (`repository/`) for data access.
    - `service/`: Business logic layer.
  - `pkg/`: Shared packages, such as database connection utilities (`pkg/db/`).
  - `seedData/`: JSON files containing sample data for various tables.
- `frontend/`: Contains the Next.js application for the user interface.
  - `src/app/`: Main application routes, layout, global styles, and page components.
    - `(components)/`: Shared UI components like Navbar, Sidebar, Header, Rating.
    - `dashboard/`, `expenses/`, `inventory/`, `products/`, `settings/`, `users/`: Feature-specific pages and components.
  - `src/state/`: Redux store configuration (`redux.tsx`), global state slices (`index.tsx`), and API query definitions (`api.tsx`).
  - `public/`: Static assets like images.

## API Endpoints (Backend Overview)

The backend exposes several API endpoints to support the frontend functionalities:

- **`GET /dashboard`**: Retrieves aggregated data for the dashboard, including popular products, sales summaries, purchase summaries, expense summaries, and expenses by category.
- **`GET /products`**: Fetches a list of all products. Supports an optional `search` query parameter to filter products by name.
- **`POST /products`**: Creates a new product. Expects a JSON body with product details (ID, name, price, rating, stock quantity).
- **`GET /expenses`**: Retrieves a summary of expenses categorized by type.
- **`GET /users`**: Fetches a list of all users.
