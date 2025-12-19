import { CharStreams, CommonTokenStream } from 'antlr4ts';
import { CalcPlusLexer } from './CalcPlusLexer';
import { CalcPlusParser } from './CalcPlusParser';
import { Calculator } from './Calculator';

describe('Calc3 Interpreter', () => {

    function runCalc(code: string, inputs: string[] = []): string[] {
        const inputStream = CharStreams.fromString(code);
        const lexer = new CalcPlusLexer(inputStream);
        const tokenStream = new CommonTokenStream(lexer);
        const parser = new CalcPlusParser(tokenStream);
        const tree = parser.calc3();

        const outputs: string[] = [];
        let inputIndex = 0;

        const mockInput = () => {
            if (inputIndex < inputs.length) {
                return inputs[inputIndex++];
            }
            return "";
        };

        const mockOutput = (msg: string) => {
            outputs.push(msg);
        };

        const visitor = new Calculator(mockInput, mockOutput);
        visitor.visit(tree);

        return outputs;
    }

    test('Hello World Output', () => {
        const code = 'write(123); write(456);';
        const result = runCalc(code);
        expect(result).toEqual(['123', '456']);
    });

    test('Echo Input', () => {
        const code = 'a = read(); write(a);';
        const result = runCalc(code, ['999']);
        expect(result).toEqual(['999']);
    });

    test('Logic Integration (Max)', () => {
        const code = `
            a = read();
            b = read();
            if (a > b) { write(a); }
            else { write(b); }
        `;
        const result = runCalc(code, ['10', '20']);
        expect(result).toEqual(['20']);
    });

    test('Invalid Input Handling', () => {
        const code = 'a = read(); write(a);';
        const result = runCalc(code, ['hello']);
        expect(result).toEqual(['0']);
    });
});