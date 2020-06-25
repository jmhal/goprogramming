package main

import (
   "log"
    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/ec2"
)

func main() {

   // Criar uma sessão para os serviços da Amazon na Região us-west-1 
   svc := ec2.New(session.New(&aws.Config{Region: aws.String("us-west-1")}))

   // Detalhes da instância a ser criada..
   runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
      ImageId:      aws.String("ami-063aa838bd7631e0b"), // Aqui tem que ter o identificador da instância do Ubuntu. Varia de região para região.
      InstanceType: aws.String("t2.micro"),     // Tamanho da instância. Vamos ficar nessa configuração.
      MinCount:     aws.Int64(1),               // Este valor e o próximo devem ficar em 1, pois vamos criar só uma.
      MaxCount:     aws.Int64(1),
   })

   // Verifica se deu erro.
   if err != nil {
       log.Fatal(err)
       return
   }

   // Avisa se conseguiu criar a instância.
   log.Println("Instância", *runResult.Instances[0].InstanceId)

   // Coloca umas tags para identificar a instância.
   _ , errtag := svc.CreateTags(&ec2.CreateTagsInput{
       Resources: []*string{runResult.Instances[0].InstanceId},
       Tags: []*ec2.Tag{
           {
               Key:   aws.String("Nome"),
               Value: aws.String("MinhaPrimeiraInstancia"),
           },
       },
   })

   if errtag != nil {
      log.Println("Impossível criar tags para instância.", runResult.Instances[0].InstanceId, errtag)
      return
   }

   log.Println("Instância identificada com sucesso")

}
