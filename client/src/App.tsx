import './App.css';
import React from 'react';
import { BrowserRouter, RouterProvider, Route, Link, redirect, createBrowserRouter } from 'react-router-dom';
import Actions from './components/Actions/Actions';
import AttackChainBuilderDnD from './components/AttackChainBuilder/AttackChainBuilderDnD';
import Action from './components/Actions/Action';
// import AttackChains from '../old_code/AttackChains/AttackChains';
// import Capabilities from './components/Capabilities/Capabilities';
// import Businesses from './components/Businesses/Businesses';
// import Exposures from './components/Exposures/Exposures';
// import Resources from './components/Resources/Resources';
// import Threats from './components/Threats/Threats';
// import Vulnerabilities from './components/Vulnerabilities/Vulnerabilities';
// import AttackChain from '../old_code/AttackChains/AttackChain';
// import Business from './components/Businesses/Business';
// import Capability from './components/Capabilities/Capability';
// import Exposure from './components/Exposures/Exposure';
// import Resource from './components/Resources/Resource';
// import Threat from './components/Threats/Threat';
// import Vulnerability from './components/Vulnerabilities/Vulnerability';


const router = createBrowserRouter([
  {
    path: "/action",
    element: <Action />
  },
    {
      path: "/actions",
      element: <Actions />
    },
    {
      path: "/attackChainBuilder",
      element: <AttackChainBuilderDnD />
    },
    // {
    //   path: "/attackChain/:id",
    //   element: <AttackChain />
    // },
    // {
    //   path: "/attackChains",
    //   element: <AttackChains />
    // },
    // {
    //   path: "/business/:id",
    //   element: <Business />
    // },
    // {
    //   path: "/businesses",
    //   element: <Businesses />
    // },
    // {
    //   path: "/capability/:id",
    //   element: <Capability />
    // },
    // {
    //   path: "/capabilities",
    //   element: <Capabilities />
    // },
    // {
    //   path: "/exposure/:id",
    //   element: <Exposure />
    // },
    // {
    //   path: "/exposures",
    //   element: <Exposures />
    // },
    // {
    //   path: "/resource/:id",
    //   element: <Resource />
    // },
    // {
    //   path: "/resources",
    //   element: <Resources />
    // },
    // {
    //   path: "/threat/:id",
    //   element: <Threat />
    // },
    // {
    //   path: "/threats",
    //   element: <Threats />
    // },
    // {
    //   path: "/vulnerability/:id",
    //   element: <Vulnerability />
    // },
    // {
    //   path: "/vulnerabilities",
    //   element: <Vulnerabilities />
    // },
  
])


const App = () => (
    <React.StrictMode>
      <RouterProvider router={router} />
    </React.StrictMode>
  );


export default App;
