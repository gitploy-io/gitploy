import './App.less'
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

import Home from './views/Home'
import Repo from './views/Repo'

function App() {
  return (
    <div className="App">
      <Router>
        <Switch>
          <Route path="/:namespace/:name">
            <Repo />
          </Route>
          <Route path="/">
            <Home />
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
