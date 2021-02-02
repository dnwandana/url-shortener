import { BrowserRouter as Router, Switch, Route } from "react-router-dom"
import Home from "./views/Home"
import loadable from "@loadable/component"

const PageNotFound = loadable(() => import("./views/PageNotFound"))

function App() {
  return (
    <div className="flex flex-col h-screen">
      <Router>
        <div className="flex-grow">
          <Switch>
            <Route path="/404">
              <PageNotFound />
            </Route>
            <Route path="/">
              <Home />
            </Route>
          </Switch>
        </div>
      </Router>
    </div>
  )
}

export default App
