Сериализатор	Десериализатор	Что преобразует
```
Стандартные		
StringSerializer	StringDeserializer	Строки в байты и обратно
IntegerSerializer	IntegerDeserializer	Целые числа (int) в байты и обратно
Int64Serializer	Int64Deserializer	Большие целые числа (long) в байты и обратно
DoubleSerializer	DoubleDeserializer	Значения с плавающей точкой (double) в байты и обратно
Специализированные		
avro.NewGenericSerializer	avro.NewGenericDeserializer	Данные в формате Avro и обратно
jsonschema.NewSerializer	jsonschema.NewDeserializer	Объекты в JSON-строки и обратно
protobuf.NewSerializer	protobuf.NewDeserializer	Работают с данными в формате Protocol Buffers (protobuf), сериализацией от Google
```
