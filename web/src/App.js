import React from 'react';
import {useEffect, useState} from 'react';
import { Routes, Route, Navigate } from "react-router-dom";
import { Layout } from 'antd';
import Calculation from "./Calculation";
import Upload from "./Upload";
import CustomLink from "./CustomLink";
import './App.css';
import 'antd/dist/antd.css';
import {HistoryInfoContext} from "./historyInfoContext";

const { Header, Content, Footer } = Layout;

function App() {
    const [error, setError] = useState(null);
    const [loading, setLoading] = useState(false);
    const {getInfo} = React.useContext(HistoryInfoContext)

    useEffect(() => {
        setLoading(true);

        getInfo().catch(err => {
            setError(err)
        })

        setLoading(false);
    }, []);

    return (
      <Layout>
        <Header style={{ position: 'fixed', zIndex: 1, width: '100%' }}>
          <div className="logo" />
            <div className="nav">
                <CustomLink to="/">History</CustomLink>
                <CustomLink to="upload">Upload</CustomLink>
            </div>
        </Header>
        <Content className="site-layout" style={{ padding: '0 50px', marginTop: 64 }}>
          <div className="site-layout-background" style={{ padding: 24, height: 'calc(100vh - 135px)' }}>
              <Routes>
                  <Route path="/" element={<Calculation />} />
                  <Route path="upload" element={<Upload />} />
                  <Route
                      path="*"
                      element={<Navigate to="/" replace />}
                  />
              </Routes>
          </div>
        </Content>
        <Footer style={{ textAlign: 'center' }}>Exchange History ©2022</Footer>
      </Layout>
  );
}

export default App;
