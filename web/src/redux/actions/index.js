import * as types from './actionTypes';

export function homepageFetchCategoriesSuccess(items) { //Action Creator
    return {
        type: types.FETCH_HP_CATEGORIES_SUCCESS,
        items
    }
};


export const getHomepageCategories = () => {
    return (dispatch) => {
        fetch('http://localhost:8083/categories')
            .then(res => res.json())
            .then(data => {
                console.dir(data.data);
                dispatch(homepageFetchCategoriesSuccess(data)); //Dispatch an action
        })
    }
}




