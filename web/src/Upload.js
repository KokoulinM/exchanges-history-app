import { LoadingOutlined, PlusOutlined } from '@ant-design/icons';
import { SmileOutlined, FrownOutlined } from '@ant-design/icons';
import { message, Upload as AntUpload, notification } from 'antd';
import React, { useState } from 'react';
import './Upload.css'
import api from "./api";

const getBase64 = (img, callback) => {
    const reader = new FileReader();
    reader.addEventListener('load', () => callback(reader.result));
    reader.readAsDataURL(img);
};

const openSuccessNotification = (placement) => {
    notification.info({
        message: `The CSV file has been uploaded successfully`,
        placement,
        icon: <SmileOutlined style={{ color: '#e91010' }} />,
    });
};

const openFailureNotification = (placement, err) => {
    notification.info({
        message: `An error occurred while uploading the file`,
        description: err,
        placement,
        icon: <FrownOutlined style={{ color: 'E91010FF' }} />,
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
    const [csvUrl, setCsvUrl] = useState();

    const uploadFile = (data) => {
        setLoading(true);
        const bodyFormData = new FormData();
        bodyFormData.append('file', data.file);

        api.post(`/history/exchanges/${data.filename}`, bodyFormData, {
            headers: {
                'content-type': 'multipart/form-data'
            }
        })
            .then(res => {
                openSuccessNotification('bottomLeft')
            })
            .catch(err => {
                openFailureNotification('bottomLeft', err)
            })
            .finally(() => {
                setLoading(false);
            })
    }

    const handleChange = (info) => {
        if (info.file.status === 'uploading') {
            return;
        }

        if (info.file.status === 'done') {
            // Get this url from response in real world.
            getBase64(info.file.originFileObj, (url) => {
                setCsvUrl(url);
            });
        }
    };

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
                onChange={handleChange}
                listType="picture-card"
                className="avatar-uploader"
                showUploadList={false}
            >
                {csvUrl ? <img src={csvUrl} alt="avatar" /> : uploadButton}
            </AntUpload>
        </div>
    )
}

export default Upload;