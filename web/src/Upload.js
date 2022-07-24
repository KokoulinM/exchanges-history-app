import { LoadingOutlined, PlusOutlined } from '@ant-design/icons';
import { SmileOutlined, FrownOutlined } from '@ant-design/icons';
import { message, Upload as AntUpload, notification } from 'antd';
import React, { useState } from 'react';
import './Upload.css'
import api from "./api";
import { HistoryInfoContext } from "./historyInfoContext";

const openSuccessNotification = (placement) => {
    notification.info({
        message: `The CSV file has been uploaded successfully`,
        placement,
        icon: <SmileOutlined style={{ color: '#1890ff' }} />,
    });
};

const openFailureNotification = (placement, err) => {
    notification.info({
        message: `An error occurred while uploading the file`,
        description: err,
        placement,
        icon: <FrownOutlined style={{ color: '#e91010' }} />,
    });
};

const beforeUpload = (file) => {
    const isCsv = file.type === 'text/csv';

    if (!isCsv) {
        message.error('You can only upload CSV file!');
    }

    const isLt2M = file.size / 1024 / 1024 < 2;

    if (!isLt2M) {
        message.error('CSV must smaller than 2MB!');
    }

    return isCsv && isLt2M;
};

function Upload() {
    const [loading, setLoading] = useState(false);

    const {getInfo} = React.useContext(HistoryInfoContext)

    const uploadFile = (data) => {
        setLoading(true);
        const bodyFormData = new FormData();
        bodyFormData.append('file', data.file);

        api.post(`/history/exchanges/${data.filename}`, bodyFormData, {
            headers: {
                'content-type': 'multipart/form-data'
            }
        })
            .then(() => {
                openSuccessNotification('bottomLeft')

                getInfo().catch(err => {
                    console.log(err)
                })
            })
            .catch(err => {
                openFailureNotification('bottomLeft', err)
            })
            .finally(() => {
                setLoading(false);
            })
    }

    const uploadButton = (
        <div>
            {loading ? <LoadingOutlined /> : <PlusOutlined />}
            <div
                style={{
                    marginTop: 8,
                }}
            >
                Upload
            </div>
        </div>
    );

    return (
        <div className="upload">
            <AntUpload
                accept="text/csv"
                customRequest={uploadFile}
                listType="picture-card"
                className="avatar-uploader"
                showUploadList={false}
                beforeUpload={beforeUpload}
            >
                {uploadButton}
            </AntUpload>
        </div>
    )
}

export default Upload;