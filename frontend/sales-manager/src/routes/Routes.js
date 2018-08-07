import React from "react";
import { Route, Switch } from "react-router-dom";
import Signup from "../container/Signup";
import AppliedRoute from "../components/AppliedRoute"

export default ({ childProps }) =>
<Switch>
    <AppliedRoute path="/signup" exact component={Signup} props={childProps} />
</Switch>
