# Frontend

This is the frontend microservice for the Kvis application. It is built using React and Vite, and it communicates with a backend service to fetch deployment data.

## Project Structure

- `src/`
  - `assets/`: Contains static assets like images.
  - `components/`: Contains reusable React components.
  - `types/`: Contains TypeScript type definitions.
  - `utils/`: Contains utility functions.
  - `App.tsx`: Main App component.
  - `Deployments.tsx`: Component for displaying deployments.
  - `main.tsx`: Entry point for the React application.
  - `index.css`: Global CSS styles.

## Build and Run Docker container

To build and run the Docker container for the frontend, use the following command:

```sh
./pipeline.sh
```

## Development

To start the development server, use the following command:

```sh
npm run dev
```

## Deployment

The `Deployments.tsx` component fetches a list of deployments from the backend service and displays them in a drawer. The drawer can be opened and closed using buttons.

## Routing

The application uses React Router for client-side routing. The main routes are defined in `main.tsx`:

- `/`: Renders the `App` component.
- `/deployments`: Renders the `Deployments` component.

## Styling

The application uses CSS for styling, with global styles defined in `index.css`. The Material-UI library is used for UI components like the AppBar and Drawer.

## Environment

The project uses Vite for development and build processes. TypeScript is used for type checking and ensuring code quality.

## Additional Information

For more details on Vite and React, click on the logos in the main page of the application.
