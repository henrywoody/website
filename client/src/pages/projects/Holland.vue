<template>
    <main>
        <h1>Holland</h1>

        <section class="description">
            <p>
                Holland is a genetic algorithm library for Python. The package is application-agnostic and can be used on any problem for which there is a way to encode solutions and evalute the success of a potential solution. The package is named for <ExternalLink href="https://en.wikipedia.org/wiki/John_Henry_Holland">John H. Holland</ExternalLink>, a pioneer in genetic algorithms research.
            </p>

            <p>
                The genetic algorithm is a search algorithm for finding solutions to problems. It is a metaheuristic, meaning that it is a heuristic for finding heuristics for solving other problems. In order for the genetic algorithm to be used on a problem, two things are required: a way to encode solutions to the problem and a method for evaluating the "fitness" of a possible solution.
            </p>

            <p>
                With these two components the general process for the genetic algorithm is:

                <ol>
                    <li>Generate initial (random) solutions</li>
                    <li>Evaluate each solution's fitness</li>
                    <li>Generate new solutions by performing crossover and mutation on existing solutions according to their fitness</li>
                    <li>Return to step 2</li>
                </ol>
            </p>

            <p>
                For such a simple algorithm, solutions found from employing the genetic algorithm can be quite complex and highly performant.
            </p>

            <p>
                Holland can be installed using pip with <code>pip install holland</code>.
            </p>

            <p>
                The repository <ExternalLink href="https://github.com/lambdalife/holland-gym">Holland-Gym</ExternalLink> contains examples of evolved solutions to some environments from <ExternalLink href="http://gym.openai.com/">Open AI's Gym</ExternalLink>.
            </p>

            <h2>Hello World</h2>

            <p>
                Here is a "Hello World!" example:
            </p>

            <CodeBlock
                language="python"
                :code="helloWorldCode"
            />

            <p>
                With a sample run:
            </p>

            <pre>
Generation: 0;    Top Score: 201:     N~flx.JGcu-*
Generation: 1;    Top Score: 98:      Xljlw);mj]f
Generation: 2;    Top Score: 64:      =c}kk SmsYf
Generation: 3;    Top Score: 37:      Kcjlk$Vms]f
Generation: 4;    Top Score: 24:      Cdjkn Smshf
Generation: 5;    Top Score: 16:      Idjln Vmshf
Generation: 6;    Top Score: 14:      Idjln Voshf
Generation: 7;    Top Score: 11:      Hdjln Vmslf
Generation: 8;    Top Score: 9:       Hdjln Voslf
Generation: 9;    Top Score: 8:       Hdjln Vosle
Generation: 10;   Top Score: 7:       Hdmln Vosle
Generation: 11;   Top Score: 6:       Hdlln Vosle
Generation: 12;   Top Score: 5:       Hdllo Vosle
Generation: 13;   Top Score: 4:       Hdllo Vosle!
Generation: 14;   Top Score: 3:       Hello Vosle!
Generation: 15;   Top Score: 2:       Hello Wosle!
Generation: 16;   Top Score: 2:       Hello Wosle!
Generation: 17;   Top Score: 1:       Hello Worle!
Generation: 18;   Top Score: 1:       Hello Worle!
Generation: 19;   Top Score: 1:       Hello Worle!
Generation: 20;   Top Score: 0:       Hello World!
            </pre>

            <p>
                Best genome:
            </p>

            <CodeBlock
                language="python"
                :code="helloWorldGenome"
            />
        </section>

        <GithubLinkIcon href="https://github.com/lambdalife/holland"/>
    </main>
</template>

<script>
import CodeBlock from "../../components/CodeBlock";
import ExternalLink from "../../components/ExternalLink";
import GithubLinkIcon from "../../components/GithubLinkIcon";

export default {
    name: "Holland",
    components: {
        CodeBlock,
        ExternalLink,
        GithubLinkIcon,
    },
    data() {
        return {
helloWorldCode: `from holland import Evolver
from holland.library import get_uniform_crossover_function
from holland.utils import bound_value
import random


# Define a fitness function
def fitness_function(genome):
    message = genome["message"]
    target = "Hello World!"
    score = 0
    for i in range(len(message)):
        score += abs(ord(target[i]) - ord(message[i]))
    return score

def mutation_function(value):
    mutated_value = ord(value) * random.random() * 2
    return chr(
        bound_value(
            mutated_value,
            minimum=32, maximum=126, to_int=True
        )
    )

# Define genome parameters for individuals
genome_params = {
    "message": {
        "type": "[str]",
        "size": len("Hello World!"),
        "initial_distribution": lambda: chr(random.randint(32, 126)),
        "crossover_function": get_uniform_crossover_function(),
        "mutation_function": mutation_function,
        "mutation_rate": 0.15
    }
}

# Define how to select individuals for reproduction
selection_strategy = {"pool": {"top": 10}}

# Run Evolution
evolver = Evolver(
    fitness_function,
    genome_params,
    selection_strategy,
    should_maximize_fitness=False
)
final_population = evolver.evolve(
    stop_conditions={"target_fitness": 0}
)`,
helloWorldGenome: `{
    'message': ['H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd', '!']
}`
        }
    }
}
</script>

<style scoped>
section {
    width: 85%;
    max-width: 30rem;
    margin: 0 auto 1rem;
}
</style>