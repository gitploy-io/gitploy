import './App.less'
import {
    BrowserRouter as Router,
    Switch,
    Route,
} from "react-router-dom";

import Home from "./views/home"
import Repo from "./views/Repo"
import Deployment from "./views/Deployment"
import Settings from "./views/Settings"
import Members from "./views/members"

function App(): JSX.Element {
    return (
        <div className="App">
            <Router>
                <Switch>
                    <Route path="/:namespace/:name/deployments/:number">
                        <Deployment />
                    </Route>
                    <Route path="/:namespace/:name/:tab">
                        <Repo />
                    </Route>
                    <Route path="/:namespace/:name">
                        <Repo />
                    </Route>
                    <Route path="/settings">
                        <Settings />
                    </Route>
                    <Route path="/members">
                        <Members />
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
