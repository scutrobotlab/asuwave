import { getVariable, getVariableType } from "@/api/variable.js";

export default {
    namespaced: true,
    state: {
        variables: {
            proj: {},
            read: {},
            modi: {},
        },
        vTypes: [],
    },
    getters: {
        searchVToProj: (state) => (keyword) => {
            if (keyword) {
                return state.variables["proj"].filter(function (product) {
                    return Object.keys(product).some(function (key) {
                        return (
                            String(product[key])
                                .toLowerCase()
                                .indexOf(keyword) > -1
                        );
                    });
                });
            } else {
                return state.variables["proj"];
            }
        },
    },
    mutations: {
        setV(state, i) {
            state.variables[i.t] = i.v;
        },
        setVType(state, i) {
            state.vTypes = i;
        },
    },
    actions: {
        getV({ commit }, t) {
            getVariable(t).then((v) => {
                commit("setV", { t, v });
            });
        },
        getVType({ commit }) {
            getVariableType().then((v) => {
                commit("setVType", v);
            });
        },
    },
    modules: {},
};