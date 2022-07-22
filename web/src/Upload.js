import { LoadingOutlined, PlusOutlined } from '@ant-design/icons';
import { message, Upload as AntUpload } from 'antd';
import React, { useState } from 'react';
import './Upload.css'

const getBase64 = (img, callback) => {
    const reader = new FileReader();
    reader.addEventListener('load', () => callback(reader.result));
    reader.readAsDataURL(img);
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

    const handleChange = (info) => {
        if (info.file.status === 'uploading') {
            setLoading(true);
            return;
        }

        if (info.file.status === 'done') {
            // Get this url from response in real world.
            getBase64(info.file.originFileObj, (url) => {
                setLoading(false);
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
                name="avatar"
                listType="picture-card"
                className="avatar-uploader"
                showUploadList={false}
                action="https://www.mocky.io/v2/5cc8019d300000980a055e76"
                beforeUpload={beforeUpload}
                onChange={handleChange}
            >
                {csvUrl ? <img src={csvUrl} alt="avatar" /> : uploadButton}
            </AntUpload>
        </div>
    )
}

export default Upload;