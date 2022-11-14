package info.sqldev.instance;

public class InstanceDto {
    int result;
    String msg;
    Instance data;

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

    public Instance getData() {
        return data;
    }

    public void setData(Instance data) {
        this.data = data;
    }
}
