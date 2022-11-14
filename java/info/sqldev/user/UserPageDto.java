package info.sqldev.user;

public class UserPageDto {
    int result;
    String msg;
    UserPage data;

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

    public UserPage getData() {
        return data;
    }

    public void setData(UserPage data) {
        this.data = data;
    }
}
