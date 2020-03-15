const initialState = {
    isAuthenticated: false,
    disable: false,
    user: null,
    token: null,
};

const reducer = (state = initialState, action) => {

    switch (action.type){
      case "LOGIN":
        return{
          ...state,
          isAuthenticated: true,
          disable: true,
          user: action.payload.user,
          token: action.payload.token
        };
      case "SIGNUP":
        return{
          ...state,
          isAuthenticated: true,
          disable: true,
          user: action.payload.user,
          token: action.payload.token
        };
        case "LOGOUT":
          localStorage.clear();
          return {
            ...state,
            isAuthenticated: false,
            user: null,
            token: null
        };
        default:
          return state
    
    }
  }

  export default reducer