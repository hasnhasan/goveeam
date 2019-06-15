#!/home/toast/anaconda3/bin/python3
from flask import Flask
from flask import Response

app = Flask(__name__)


@app.route('/api', methods=['GET'])
def index():
    result = """
    <EnterpriseManager xmlns="http://www.veeam.com/ent/v1.0" xmlns:xsd="http://www.w3.org/2001/XMLSchema" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
 <Links>
   <Link Href="http://localhost:9399/api/logonSessions" Type="LogonSessionList" Rel="Down"/>
   <Link Href="http://localhost:9399/api/sessionMngr/?v=latest" Type="LogonSession" Rel="Create"/>
 </Links>
 <SupportedVersions>
   <SupportedVersion Name="v1">
     <Links>
       <Link Href="http://localhost:9399/api/sessionMngr/?v=v1" Type="LogonSession" Rel="Create"/>
     </Links>
   </SupportedVersion>
   <SupportedVersion Name="v1_1">
     <Links>
       <Link Href="http://localhost:9399/api/sessionMngr/?v=v1_1" Type="LogonSession" Rel="Create"/>
     </Links>
   </SupportedVersion>
   <SupportedVersion Name="v1_2">
     <Links>
       <Link Href="http://localhost:9399/api/sessionMngr/?v=v1_2" Type="LogonSession" Rel="Create"/>
     </Links>
   </SupportedVersion>
   <SupportedVersion Name="v1_3">
     <Links>
       <Link Href="http://localhost:9399/api/sessionMngr/?v=v1_3" Type="LogonSession" Rel="Create"/>
     </Links>
   </SupportedVersion>
   <SupportedVersion Name="v1_4">
     <Links>
       <Link Href="http://localhost:9399/api/sessionMngr/?v=v1_4" Type="LogonSession" Rel="Create"/>
     </Links>
   </SupportedVersion>
 </SupportedVersions>
</EnterpriseManager>
    """
    return Response(result, mimetype="application/xml", status=200)


@app.route('/api/sessionMngr', methods=['GET','POST'])
def login():
    result = """
        <LogonSession xmlns="http://www.veeam.com/ent/v1.0" Type="LogonSession" Href="http://localhost:9399/api/logonSessions/5496707f-c814-47ce-8d6d-110aa03cec03">
 <Links>
   <Link Rel="Up" Type="EnterpriseManager" Href="http://localhost:9399/api/" />
   <Link Rel="Down" Type="BackupServerReferenceList" Href="http://localhost:9399/api/backupServers" />
   <Link Rel="Down" Type="ManagedServerReferenceList" Href="http://localhost:9399/api/managedServers" />
   <Link Rel="Down" Type="JobReferenceList" Href="http://localhost:9399/api/jobs" />
   <Link Rel="Down" Type="FailoverPlanReferenceList" Href="http://localhost:9399/api/failoverPlans" />
   <Link Rel="Down" Type="HierarchyRootReferenceList" Href="http://localhost:9399/api/hierarchyRoots" />
   <Link Rel="Down" Type="RepositoryReferenceList" Href="http://localhost:9399/api/repositories" />
   <Link Rel="Down" Type="BackupReferenceList" Href="http://localhost:9399/api/backups" />
   <Link Rel="Down" Type="RestorePointReferenceList" Href="http://localhost:9399/api/restorePoints" />
   <Link Rel="Down" Type="VmRestorePointReferenceList" Href="http://localhost:9399/api/vmRestorePoints" />
   <Link Rel="Down" Type="VAppRestorePointReferenceList" Href="http://localhost:9399/api/vAppRestorePoints" />
   <Link Rel="Down" Type="ReplicaReferenceList" Href="http://localhost:9399/api/replicas" />
   <Link Rel="Down" Type="VmReplicaPointReferenceList" Href="http://localhost:9399/api/vmReplicaPoints" />
   <Link Rel="Down" Type="CatalogVmReferenceList" Href="http://localhost:9399/api/catalog/vms" />
   <Link Rel="Down" Type="BackupJobSessionReferenceList" Href="http://localhost:9399/api/backupSessions" />
   <Link Rel="Down" Type="RestoreSessionReferenceList" Href="http://localhost:9399/api/restoreSessions" />
   <Link Rel="Down" Type="ReplicaJobSessionReferenceList" Href="http://localhost:9399/api/replicaSessions" />
   <Link Rel="Down" Type="BackupTaskSessionReferenceList" Href="http://localhost:9399/api/backupTaskSessions" />
   <Link Rel="Down" Type="ReplicaTaskSessionReferenceList" Href="http://localhost:9399/api/replicaTaskSessions" />
   <Link Rel="Down" Type="EnterpriseSecuritySettings" Href="http://localhost:9399/api/security" />
   <Link Rel="Down" Type="WanAcceleratorReferenceList" Href="http://localhost:9399/api/wanAccelerators" />
   <Link Rel="Down" Type="BackupFileReferenceList" Href="http://localhost:9399/api/backupFiles" />
   <Link Rel="Down" Type="TaskList" Href="http://localhost:9399/api/tasks" />
   <Link Rel="Down" Type="QueryService" Href="http://localhost:9399/api/querySvc" />
   <Link Rel="Down" Type="LookupService" Href="http://localhost:9399/api/lookupSvc" />
   <Link Rel="Down" Type="Report" Href="http://localhost:9399/api/reports/summary" Name="Summary" />
   <Link Rel="Down" Type="BackupServerList" Href="http://localhost:9399/api/backupServers?format=Entity" />
   <Link Rel="Down" Type="ManagedServerList" Href="http://localhost:9399/api/managedServers?format=Entity" />
   <Link Rel="Create" Href="http://localhost:9399/api/backupServers?action=create" />
   <Link Rel="Down" Type="JobList" Href="http://localhost:9399/api/jobs?format=Entity" />
   <Link Rel="Down" Type="FailoverPlanList" Href="http://localhost:9399/api/failoverPlans?format=Entity" />
   <Link Rel="Down" Type="HierarchyRootList" Href="http://localhost:9399/api/hierarchyRoots?format=Entity" />
   <Link Rel="Down" Type="RepositoryList" Href="http://localhost:9399/api/repositories?format=Entity" />
   <Link Rel="Down" Type="BackupList" Href="http://localhost:9399/api/backups?format=Entity" />
   <Link Rel="Down" Type="RestorePointList" Href="http://localhost:9399/api/restorePoints?format=Entity" />
   <Link Rel="Down" Type="VmRestorePointList" Href="http://localhost:9399/api/vmRestorePoints?format=Entity" />
   <Link Rel="Down" Type="VAppRestorePointList" Href="http://localhost:9399/api/vAppRestorePoints?format=Entity" />
   <Link Rel="Down" Type="ReplicaList" Href="http://localhost:9399/api/replicas?format=Entity" />
   <Link Rel="Down" Type="VmReplicaPointList" Href="http://localhost:9399/api/vmReplicaPoints?format=Entity" />
   <Link Rel="Down" Type="CatalogVmList" Href="http://localhost:9399/api/catalog/vms?format=Entity" />
   <Link Rel="Down" Type="BackupJobSessionList" Href="http://localhost:9399/api/backupSessions?format=Entity" />
   <Link Rel="Down" Type="RestoreSessionList" Href="http://localhost:9399/api/restoreSessions?format=Entity" />
   <Link Rel="Down" Type="ReplicaJobSessionList" Href="http://localhost:9399/api/replicaSessions?format=Entity" />
   <Link Rel="Down" Type="BackupTaskSessionList" Href="http://localhost:9399/api/backupTaskSessions?format=Entity" />
   <Link Rel="Down" Type="ReplicaTaskSessionList" Href="http://localhost:9399/api/replicaTaskSessions?format=Entity" />
   <Link Rel="Down" Type="WanAcceleratorList" Href="http://localhost:9399/api/wanAccelerators?format=Entity" />
   <Link Rel="Down" Type="VCloudService" Href="http://localhost:9399/api/vCloud" />
   <Link Rel="Down" Type="BackupFileList" Href="http://localhost:9399/api/backupFiles?format=Entity" />
   <Link Rel="Down" Type="CloudConnectService" Href="http://localhost:9399/api/cloud" />
   <Link Rel="Delete" Href="http://localhost:9399/api/logonSessions/5496707f-c814-47ce-8d6d-110aa03cec03" />
 </Links>
 <UserName>SRV13\Administrator</UserName>
 <SessionId>5496707f-c814-47ce-8d6d-110aa03cec03</SessionId>
</LogonSession>
        """

    return Response(result, mimetype="application/xml", headers={"X-RestSvcSessionId": "O9Msj3rq7EGUhtSMBQx+mw=="},
                    status=200)


if __name__ == '__main__':
    app.run(debug=True)
