package info.sqldev.sqlquery;

public class SqlQueryDto {
    int result;
    String msg;
    SqlQuery data;

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

    public SqlQuery getData() {
        return data;
    }

    public void setData(SqlQuery data) {
        this.data = data;
    }
}

