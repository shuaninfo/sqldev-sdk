package info.sqldev.sqlquery;

public class SqlQueryForm {
    private long project;
    private String iid;
    private String sql;
    private String owner;
    private String schema;
    private int pageIndex;
    private int pageSize;
    private boolean needTotal;

    public long getProject() {
        return project;
    }

    public void setProject(long project) {
        this.project = project;
    }

    public String getIid() {
        return iid;
    }

    public void setIid(String iid) {
        this.iid = iid;
    }

    public String getSql() {
        return sql;
    }

    public void setSql(String sql) {
        this.sql = sql;
    }

    public String getOwner() {
        return owner;
    }

    public void setOwner(String owner) {
        this.owner = owner;
    }

    public String getSchema() {
        return schema;
    }

    public void setSchema(String schema) {
        this.schema = schema;
    }

    public int getPageIndex() {
        return pageIndex;
    }

    public void setPageIndex(int pageIndex) {
        this.pageIndex = pageIndex;
    }

    public int getPageSize() {
        return pageSize;
    }

    public void setPageSize(int pageSize) {
        this.pageSize = pageSize;
    }

    public boolean isNeedTotal() {
        return needTotal;
    }

    public void setNeedTotal(boolean needTotal) {
        this.needTotal = needTotal;
    }
}
