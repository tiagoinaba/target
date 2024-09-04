`go run .` para executar as soluções

# 1 - Fibonacci
Defini uma função que calcula a sequência de Fibonacci enquanto o número atual da sequência é menor que o número inserido pelo usuário. Quando o número da sequência for maior ou igual, paro o loop e comparo o o último número da sequência com o número inserido. Se forem iguais, o número pertence à sequência.

# 2 - String contém 'a' ou 'A'
Defino uma variável `count` que é incrementada toda vez que encontro um dos dois caracteres na string inserida.

# 3 - Ao final do processamento, qual será o valor da variável SOMA? 
Resposta: 77

# 4 - Descubra a lógica
- a) 9
- b) 128
- c) 49
- d) 100
- e) 13
- f) 20

# 5 - Como você faria para descobrir, usando apenas duas idas até uma das salas das lâmpadas, qual interruptor controla cada lâmpada?
Tem uma série que chama *Alice in Borderland* que expõe este problema, mas a solução consiste em acender um dos interruptores e deixá-lo aceso por alguns minutos. Após isso, apagá-lo e acender algum dos outros dois. Em seguida, utilizamos a primeira visita à uma das salas. Se:

- a luz estiver acesa, o interruptor que ligamos por último controla esta lâmpada.
- a luz estiver apagada mas a lâmpada está quente, o primeiro interruptor que acendemos controla esta lâmpada.
- a luz estiver apagada e gelada, o interruptor que não ligamos controla esta lâmpada.

Depois, é só visitar alguma das outras salas e fazer as mesmas verificações. Sabendo quais interruptores controlam duas das lâmpadas, podemos deduzir qual o último interruptor controla.
