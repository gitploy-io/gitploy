import './App.less'
import {
  BrowserRouter as Router,
  Switch,
  Route,
} from "react-router-dom";

import Home from './views/Home'

function App() {
  return (
    <div className="App">
      <Router>
        <Switch>
          <Route path="/">
            <Home />
          </Route>
          <Route path="/:namespace/:name">
            <Home />
          </Route>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
