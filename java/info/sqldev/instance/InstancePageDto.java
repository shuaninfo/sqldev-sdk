package info.sqldev.instance;

public class InstancePageDto {
    int result;
    String msg;
    InstancePage data;

    public int getResult() {
        return result;
    }

    public void setResult(int result) {
        this.result = result;
    }

    public String getMsg() {
        return msg;
    }

    public void setMsg(String msg) {
        this.msg = msg;
    }

    public InstancePage getData() {
        return data;
    }

    public void setData(InstancePage data) {
        this.data = data;
    }
}
