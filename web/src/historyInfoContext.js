import React from "react";
import api from "./api";

export const HistoryInfoContext = React.createContext(null);

const HistoryInfoProvider = ({ children }) => {
    const [cryptoCurrencies, setCryptoCurrencies] = React.useState([])
    const [payMethods, setPayMethods] = React.useState([])

    const getInfo = () => {
        return api.get('/history/info')
            .then(res => {
                setCryptoCurrencies(res.data.cryptoCurrencies)
                setPayMethods(res.data.payMethods)
            })
    }

    return <HistoryInfoContext.Provider value={{ cryptoCurrencies, payMethods, getInfo }}>{children}</HistoryInfoContext.Provider>;
};

export default HistoryInfoProvider;