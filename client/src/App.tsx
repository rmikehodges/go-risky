import './App.css';
import React from 'react';
import { RouterProvider, createBrowserRouter } from 'react-router-dom';

import AttackChainBuilderDnD from './components/AttackChainBuilder/AttackChainBuilderDnD';
import ImpactBuilder from './components/ImpactBuilder/ImpactBuilder';


import GetAction from './components/Actions/GetAction';
import ListActions from './components/Actions/ListActions';
import GetAsset from './components/Assets/GetAsset';
import ListAssets from './components/Assets/ListAssets';
import GetThreat from './components/Threats/GetThreat';
import ListThreats from './components/Threats/ListThreats';
import GetVulnerability from './components/Vulnerabilities/GetVulnerability';
import ListVulnerabilities from './components/Vulnerabilities/ListVulnerabilities';
import GetBusiness from './components/Businesses/GetBusiness';
import ListBusinesses from './components/Businesses/ListBusinesses';
import GetCapability from './components/Capabilities/GetCapability';
import ListCapabilities from './components/Capabilities/ListCapabilities';
import GetResource from './components/Resources/GetResource';
import ListResources from './components/Resources/ListResources';
import GetAttackChain from './components/AttackChains/GetAttackChain';
import ListAttackChains from './components/AttackChains/ListAttackChains';
import ListAttackChainSteps from './components/AttackChainSteps/ListAttackChainSteps';





const router = createBrowserRouter([
    {
      path: "/attackChainBuilder",
      element: <AttackChainBuilderDnD />
    },
    {
      path: "/impactBuilder",
      element: <ImpactBuilder />
    },
    {

      path: "/action",
      element: <GetAction />
    },
    {
      path: "/actions",
      element: <ListActions />
    },

    {
      path: "/asset",
      element: <GetAsset />
    },
    {
      path: "/assets",
      element: <ListAssets />
    },
    {
      path: "/threat",
      element: <GetThreat />
    },
    {
      path: "/threats",
      element: <ListThreats />
    },
    {
      path: "/vulnerability",
      element: <GetVulnerability />
    },
    {
      path: "/vulnerabilities",
      element: <ListVulnerabilities />
    },
    {
      path: "/business",
      element: <GetBusiness />
    },
    {
      path: "/businesses",
      element: <ListBusinesses />
    },
    {
      path: "/capability",
      element: <GetCapability />
    },
    {
      path: "/capabilities",
      element: <ListCapabilities />
    },
    {
      path: "/resource",
      element: <GetResource />
    },
    {
      path: "/resources",
      element: <ListResources />
    },
    {
      path: "/attackChain",
      element: <GetAttackChain />
    },
    {
      path: "/attackChains",
      element: <ListAttackChains />
    },
    {
      path: "/attackChainSteps",
      element: <ListAttackChainSteps />
    },
  
])


const App = () => (
    <React.StrictMode>
      <RouterProvider router={router} />
    </React.StrictMode>
  );


export default App;
