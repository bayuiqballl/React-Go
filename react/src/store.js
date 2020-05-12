import { createStore, applyMiddleware, compose } from "redux";
import thunk from "redux-thunk";

import siswaReducer from "./reducer/siswaReducer";

export default createStore(siswaReducer, applyMiddleware(thunk));
