import TestEdit from "@/views/admin/test/edit"


const TetsEditPage = () => <TestEdit />

TetsEditPage.acl = {
    action: 'read',
    permission: 'admin'
}
TetsEditPage.admin = true
export default TetsEditPage