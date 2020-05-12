import {
  GET_SISWA_SUCCESS,
  GET_SISWA_FAILED,
  GET_SISWA_BEGIN,
} from "../action/siswaAction";

const initialState = {
  data: [],
  error: null,
};

const siswaReducer = (state = initialState, action) => {
  switch (action.type) {
    case GET_SISWA_BEGIN:
      return {
        ...state,
      };
    case GET_SISWA_SUCCESS:
      return {
        ...state,
        data: action.result,
      };
    case GET_SISWA_FAILED:
      return {
        ...state,
        error: action.error,
      };
    default:
      return state;
  }
};

export default siswaReducer;
