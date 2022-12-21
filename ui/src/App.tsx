import './App.less';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

import Home from './views/home';
import Repo from './views/repo';
import Deployment from './views/deployment';
import Settings from './views/settings';
import Members from './views/members';
import Activities from './views/activities';

function App(): JSX.Element {
  return (
    <div className="App">
      <Router>
        <Routes>
          <Route
            path="/:namespace/:name/deployments/:number"
            element={<Deployment />}
          />
          <Route path="/:namespace/:name/:tab" element={<Repo />} />
          <Route path="/:namespace/:name" element={<Repo />} />
          <Route path="/settings" element={<Settings />} />
          <Route path="/members" element={<Members />} />
          <Route path="/activities" element={<Activities />} />
          <Route path="/" element={<Home />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
