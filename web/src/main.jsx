import React from 'react'
import ReactDOM from 'react-dom/client'
import {
  createBrowserRouter,
  RouterProvider,
} from "react-router-dom";
import './index.css'
import Root from "./routes/root";
import About from './pages/About';
import Index from './pages/Index';
import { loader as indexLoader } from './helpers/loader_index';
import Contact from './routes/contact';

const router = createBrowserRouter([
  {
    path: "/",
    element: <Root />,
    children: [
      {
        path: "contacts/:contactId",
        element: <Contact />,
      },
      {
        index: true,
        element: <Index />,
        loader: indexLoader,
      },
      {
        path: "about",
        element: <About />,
      },
    ],
  },
]);


ReactDOM.createRoot(document.getElementById('root')).render(
  <React.StrictMode>
    <RouterProvider router={router} />
  </React.StrictMode>,
)
