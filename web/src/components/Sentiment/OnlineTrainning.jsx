import { Upload, message, Card } from 'antd';
import { InboxOutlined } from '@ant-design/icons';

const { Dragger } = Upload;

export default class OnlineTrainning extends React.Component {

    constructor(props) {
        super(props);
        this.state = {
            result: null
        }

        this.onPredict = this.onPredict.bind(this)
    }

    onPredict(sentence) {
        // console.log(sentence);
        fetch("http://192.168.0.102:8080/predict?sentence="+sentence)
            .then(response => response.json())
            .then(data => {
                this.setState({result: data.sentiment})
            });
    }

    onChange(info) {
        const { status } = info.file;
        if (status !== 'uploading') {
          console.log(info.file, info.fileList);
        }
        if (status === 'done') {
          message.success(`${info.file.name} file uploaded successfully.`);
        } else if (status === 'error') {
          message.error(`${info.file.name} file upload failed.`);
        }
    }

    render() {
        return (
            <Card>
                <Dragger name="file" action="" onChange={this.onChange} >
                    <p className="ant-upload-drag-icon">
                        <InboxOutlined />
                    </p>
                    <p className="ant-upload-text">Click or drag file to this area to upload</p>
                    <p className="ant-upload-hint">
                    Support for a single or bulk upload. Strictly prohibit from uploading company data or other
                    band files
                    </p>
                </Dragger>
            </Card>
        )
    }
}