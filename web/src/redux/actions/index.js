import * as types from './actionTypes';

export function homepageFetchCategoriesSuccess(items) { //Action Creator
    return {
        type: types.FETCH_HP_CATEGORIES_SUCCESS,
        items
    }
};


export function homepageCatHasErrored(bool) {
    return {
        type: types.HP_LOADING_HAS_ERRORED,
        hasErrored: bool
    };
}

export function homepageCatIsLoading(bool) {
    return {
        type: types.HP_CAT_IS_LOADING,
        isLoading: bool
    };
}


export function fetchHpCategories(url) {
    return (dispatch) => {
        dispatch(homepageCatIsLoading(true));
        fetch(url)
            .then((response) => {
                if(!response.ok) {
                    throw Error(response.statusText);
                }
                dispatch(homepageCatIsLoading(false));
                return response;
            })
            .then((response) => response.json())
            .then((items) => dispatch(homepageFetchCategoriesSuccess(items.data)))
            .catch(() => dispatch(homepageCatHasErrored(true)));
    }
}

export const getHomepageCategories = () => {
    return (dispatch) => {
        fetch('http://localhost:8083/categories')
            .then(res => res.json())
            .then(data => {
                console.dir(data.data);
                dispatch(homepageFetchCategoriesSuccess(data.data)); //Dispatch an action
        })
    }
}




