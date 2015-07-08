{{define "Subject"}}
{{.NOTIFICATIONTYPE}} - {{.HOSTDISPLAYNAME}} - {{.SERVICEDISPLAYNAME}} is {{.HOSTSTATE}}
{{end}}

{{define "Content"}}
***** Icinga  *****

Notification Type: {{.NOTIFICATIONTYPE}}

Service: {{.SERVICEDESC}}
Host: {{.HOSTALIAS}}
Address: {{.HOSTADDRESS}}
State: {{.SERVICESTATE}}

Date/Time: {{.LONGDATETIME}}

Additional Info: {{.SERVICEOUTPUT}}

Comment: [{{.NOTIFICATIONAUTHORNAME}}] {{.NOTIFICATIONCOMMENT}}
{{end}}

{{define "Plain"}}
***** Icinga  *****

Notification Type: {{.NOTIFICATIONTYPE}}

Service: {{.SERVICEDESC}}
Host: {{.HOSTALIAS}}
Address: {{.HOSTADDRESS}}
State: {{.SERVICESTATE}}

Date/Time: {{.LONGDATETIME}}

Additional Info: {{.SERVICEOUTPUT}}

Comment: [{{.NOTIFICATIONAUTHORNAME}}] {{.NOTIFICATIONCOMMENT}}
{{end}}
