package main
import(
  "fmt"
  "github.com/aws/aws-sdk-go/service/ec2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
  "os"
  "errors"
  "log"
)

func main(){
  fmt.Println("This will pull info for all instances and check for security groups in use.")
  securityGroupIDsMap:=make(map[string]int, 0)
  securityGroupsMap:=make(map[string]*ec2.GroupIdentifier, 0)
  securityGroupsToBeDeletedMap:=make(map[string]*ec2.SecurityGroup, 0)
  log.Println("Creating aws session...")
  session,err:=GetAWSSession()
  if err!=nil{
    log.Println(err)
    os.Exit(1)
  }
  log.Println("OK")
  log.Println("Creating service")
  ec2Service:=ec2.New(session)
  if ec2Service==nil{
    log.Println("can't create ec2 service")
    os.Exit(1)
  }
  log.Println("Creating service done.")
  log.Println("Pulling instnaces data, please wait...")
  output,err:=ec2Service.DescribeInstances(&ec2.DescribeInstancesInput{
    DryRun: aws.Bool(false),
    MaxResults: aws.Int64(230),
  })
  if err!=nil{
    log.Println(err)
    os.Exit(1)
  }
  log.Println("Done")
  fmt.Println("")
  fmt.Println("here we go:")
  fmt.Println("=================================")
  totalInstances:=0
  for _,reservation:=range (output.Reservations){
    for _,instanceInfo:=range(reservation.Instances){
      for _,currentGroup:=range(instanceInfo.SecurityGroups){
        securityGroupsMap[*currentGroup.GroupId]=currentGroup
        securityGroupIDsMap[*currentGroup.GroupId]=securityGroupIDsMap[*currentGroup.GroupId]+1
      }
    }
  }
  for key,value:=range(securityGroupIDsMap){
    fmt.Printf("%s(%s) has %d instances\r\n",key,*securityGroupsMap[key].GroupName,value)
    totalInstances+=value;
  }
  fmt.Println("====================================")
  fmt.Printf("total instances: %d\r\n",totalInstances)
  fmt.Printf("total security groups in use: %d\r\n",len(securityGroupIDsMap))
  fmt.Println("\r\n")
  log.Println("getting all security groups")
  allSecurityGroups,err:=ec2Service.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{
    DryRun:aws.Bool(false),
    MaxResults:aws.Int64(200),
  })
  if err!=nil{
    log.Println(err)
    os.Exit(1)
  }
  fmt.Println("checking for empty security groups")
  for _,currentSecurityGroup:=range (allSecurityGroups.SecurityGroups){
    if securityGroupIDsMap[*currentSecurityGroup.GroupId]<1{
      securityGroupsToBeDeletedMap[*currentSecurityGroup.GroupId]=currentSecurityGroup
    }
  }
  fmt.Printf("total security groups to be deleted: %d\r\n",len(securityGroupsToBeDeletedMap))
  fmt.Printf("total security groups in east: %d\r\n",len(allSecurityGroups.SecurityGroups))
  for _,securityGroupToBeDeleted:=range(securityGroupsToBeDeletedMap){
    if securityGroupToBeDeleted.VpcId!=nil{
      fmt.Printf("(%s) %s from %s\r\n",*securityGroupToBeDeleted.GroupId, *securityGroupToBeDeleted.GroupName,*securityGroupToBeDeleted.VpcId)
    }else{
      fmt.Printf("(%s) %s - no vpc\r\n",*securityGroupToBeDeleted.GroupId,*securityGroupToBeDeleted.GroupName)
    }
    fmt.Println()
    log.Printf("deleting (%s)%s:\r\n",*securityGroupToBeDeleted.GroupId,*securityGroupToBeDeleted.GroupName)
    fmt.Println()
    _,err=ec2Service.DeleteSecurityGroup(&ec2.DeleteSecurityGroupInput{
      DryRun:aws.Bool(false),
      GroupId:securityGroupToBeDeleted.GroupId,
    })
    if err!=nil{
      log.Println(err)
    }
  }

  log.Println("All done.")
}

func GetAWSSession() (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	if err != nil {
		return nil, err
	}
	if sess == nil {
		return nil, errors.New("can't create aws session")
	}
	return sess, nil
}
