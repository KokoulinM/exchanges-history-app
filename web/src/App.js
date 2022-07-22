import { Layout, Menu } from 'antd';
import { Routes, Route, Link } from "react-router-dom";
import History from "./History";
import Upload from "./Upload";
import './App.css';
import 'antd/dist/antd.css';

const { Header, Content, Footer } = Layout;

function App() {
  return (
      <Layout>
        <Header style={{ position: 'fixed', zIndex: 1, width: '100%' }}>
          <div className="logo" />
          <Menu
              theme="dark"
              mode="horizontal"
              defaultSelectedKeys={['2']}
              items={[{
                key: 1,
                label: (
                    <Link to="/">History</Link>
                )
              },
              {
                  key: 2,
                  label: (
                      <Link to="upload">Upload</Link>
                  )
              }]}
          />
        </Header>
        <Content className="site-layout" style={{ padding: '0 50px', marginTop: 64 }}>
          <div className="site-layout-background" style={{ padding: 24, height: 'calc(100vh - 135px)' }}>
              <Routes>
                  <Route path="/" element={<History />} />
                  <Route path="upload" element={<Upload />} />
              </Routes>
          </div>
        </Content>
        <Footer style={{ textAlign: 'center' }}>Exchange History ©2022</Footer>
      </Layout>
  );
}

export default App;
